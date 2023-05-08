package offers

import (
	"policy/utils/utils"
)

type Offers struct {
	Id               string         `json:"id" `
	IsActive         *utils.BitBool `json:"is_active"`
	CreatedDateTime  *string        `json:"created_date_time"`
	ModifiedDateTime *string        `json:"modified_date_time"`
	ModifiedBy       *string        `json:"modified_by"`
	CreatedBy        *string        `json:"created_by"`
	Name             *string        `json:"name"`
	OfferCode        *string        `json:"offer_code"`
	Description      *string        `json:"description"`
	ImageUrl         *string        `json:"image_url"`
}

type OffersSave struct {
	Name        *string `json:"name"`
	OfferCode   *string `json:"offer_code"`
	Description *string `json:"description"`
	ImageUrl    *string `json:"image_url"`
	CreatedBy   *string `json:"created_by"`
}
