package admin

import (
	"policy/utils/utils"
)

type UserData struct {
	Id               string         `json:"id" `
	IsActive         *utils.BitBool `json:"is_active"`
	CreatedDateTime  *string        `json:"created_date_time"`
	ModifiedDateTime *string        `json:"modified_date_time"`
	ModifiedBy       *string        `json:"modified_by"`
	CreatedBy        *string        `json:"created_by"`
	Name             *string        `json:"name"`
	MobileNo         *string        `json:"mobile_no"`
	Email            *string        `json:"email"`
}

type UserDataSave struct {
	Name      *string `json:"name"`
	MobileNo  *string `json:"mobile_no"`
	Email     *string `json:"email"`
	CreatedBy *string `json:"created_by"`
}
