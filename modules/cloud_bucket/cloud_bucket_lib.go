package cloud_bucket

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/http"
	"net/url"
)

// HandleFileUploadToBucketLib uploads file to bucket
func HandleFileUploadToBucketLib(c *gin.Context) (*string, *error) {
	bucket := "cry_app_red_storage_bucket" //your bucket name
	var err error
	ctx := appengine.NewContext(c.Request)
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("cstorage_keys.json"))
	if err != nil {
		return nil, &err
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		return nil, &err
	}
	defer f.Close()
	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		return nil, &err
	}

	if err := sw.Close(); err != nil {

		return nil, &err
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return nil, &err
	}
	path := u.EscapedPath()
	return &path, nil

}
