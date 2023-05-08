package cloud_bucket

import (
	"cloud.google.com/go/storage"
	"policy/utils/api_response"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"io"
	"net/http"
	"net/url"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
// @Tags Upload file
// @Summary		Upload file
// @Description	Upload file
// @ID				file.upload
// @Accept			multipart/form-data
// @Produce		json
// @Param			file	formData	file			true	"this is for file upload"
// @Success		200		{string}	string			"ok"
// @Router			/file_upload [post]
func HandleFileUploadToBucket(c *gin.Context) {
	bucket := "policy" //your bucket name
	var err error
	ctx := appengine.NewContext(c.Request)
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("cstorage_keys.json"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
	defer f.Close()
	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}
	result := gin.H{
		"message":  "file uploaded successfully",
		"pathname": u.EscapedPath(),
	}

	apiResponse, respErr := json.Marshal(result)
	if respErr != nil {
		c.JSON(http.StatusOK, respErr.Error())
		return
	}
	c.JSON(http.StatusOK, api_response.Data(apiResponse))
}
