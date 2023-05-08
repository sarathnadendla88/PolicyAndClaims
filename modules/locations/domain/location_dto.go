package admin

import (
	"policy/utils/utils"
)

type Location struct {
	Id               string         `json:"id" `
	IsActive         *utils.BitBool `json:"is_active"`
	CreatedDateTime  *string        `json:"created_date_time"`
	ModifiedDateTime *string        `json:"modified_date_time"`
	ModifiedBy       *string        `json:"modified_by"`
	CreatedBy        *string        `json:"created_by"`
	CompanyName      *string        `json:"company_name"`
	Address          *string        `json:"address"`
	Latitude         *string        `json:"latitude"`
	Longitude        *string        `json:"longitude"`
	Description      *string        `json:"description"`
}

type LocationSave struct {
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Latitude    *string `json:"latitude"`
	Longitude   *string `json:"longitude"`
	Description *string `json:"description"`
	CreatedBy   *string `json:"created_by"`
}
