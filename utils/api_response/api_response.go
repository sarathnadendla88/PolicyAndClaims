package api_response

import (
	"policy/utils/constants"
	"encoding/json"
	"net/http"
)

type APIResponseModel struct {
	Message       string          `json:"message,omitempty"`
	Status        int             `json:"status,omitempty"`
	MediaBasePath string          `json:"media_base_path"`
	Body          json.RawMessage `json:"body,omitempty"`
}

func Data(obj []byte) *APIResponseModel {

	apiResponse := APIResponseModel{
		Message:       "OK",
		Status:        http.StatusOK,
		MediaBasePath: constants.CloudStorageBaseURL,
		Body:          obj,
	}
	return &apiResponse
}
