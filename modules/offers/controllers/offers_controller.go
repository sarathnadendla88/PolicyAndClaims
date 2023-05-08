package controller_offers

import (
	"encoding/json"
	"net/http"
	offers "policy/modules/offers/domain"
	offer_service "policy/modules/offers/services"
	"policy/utils/api_response"
	"policy/utils/errors"

	"github.com/gin-gonic/gin"
)

// GetOffersById
// @Summary Get Offers by Id
// @Produce  json
// @Tags Offers
// @Param offer_id path string true "Offer Id"
// @Success 200 {object} string "ok"
// @Router /offers/{offer_id} [get]
func GetOffersById(c *gin.Context) {
	offerId := c.Param("offer_id")
	offerObject, err := offer_service.OfferService.GetOfferById(offerId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(offerObject)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// GetAllOffers
// @Summary Get all offers
// @Produce  json
// @Tags Offers
// @Success 200 {object} string "ok"
// @Router /offers [get]
func GetAllOffers(c *gin.Context) {
	offers, err := offer_service.OfferService.GetAllOffers()
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	apiResponse, respErr := json.Marshal(offers)
	if respErr != nil {
		c.JSON(http.StatusInternalServerError, respErr.Error())
		return
	}

	c.JSON(http.StatusOK, api_response.Data(apiResponse))
	return
}

// AddOffer
// @Summary Add Offer
// @Produce  json
// @Tags Offers
// @Param offer body offers.OffersSave true "Offers Data"
// @Success 200 {object} offers.OffersSave
// @Router /offers [post]
func AddOffer(c *gin.Context) {
	// read json body
	var offerObject offers.Offers
	if err := c.ShouldBindJSON(&offerObject); err != nil {
		restError := errors.NewBadRequest(err.Error())
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := offer_service.OfferService.CreateOffer(offerObject)
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
