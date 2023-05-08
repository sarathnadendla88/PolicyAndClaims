package controller_location

import (
	"encoding/json"
	"net/http"
	location "policy/modules/locations/domain"
	location_service "policy/modules/locations/services"
	"policy/utils/api_response"
	"policy/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetLocationsById
// @Summary Get Locations by Id
// @Produce  json
// @Tags Locations
// @Param location_id path string true "Location Id"
// @Success 200 {object} string "ok"
// @Router /locations/{location_id} [get]
func GetLocationById(c *gin.Context) {
	locationId := c.Param("location_id")
	locationObj, err := location_service.LocationService.GetLocationById(locationId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(locationObj)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// GetAllLocations
// @Summary Get all locations
// @Produce  json
// @Tags Locations
// @Success 200 {object} string "ok"
// @Router /locations [get]
func GetAllLocations(c *gin.Context) {
	location, err := location_service.LocationService.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(location)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// AddLocation
// @Summary Add Location
// @Produce  json
// @Tags Locations
// @Param location body location.LocationSave true "Locations Data"
// @Success 200 {object} location.LocationSave
// @Router /locations [post]
func AddLocation(c *gin.Context) {
	// read json body
	var locationObject location.Location
	if err := c.ShouldBindJSON(&locationObject); err != nil {
		restError := errors.NewBadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := location_service.LocationService.CreateLocation(locationObject)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	apiResponse, respErr := json.Marshal(result)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
}
