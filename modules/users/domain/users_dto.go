package user

import (
	"policy/utils/utils"
)

type User struct {
	Id               string         `json:"id" `
	IsActive         *utils.BitBool `json:"is_active"`
	CreatedDateTime  *string        `json:"created_date_time"`
	ModifiedDateTime *string        `json:"modified_date_time"`
	ModifiedBy       *string        `json:"modified_by"`
	CreatedBy        *string        `json:"created_by"`
	RegisteredDate   *string        `json:"registered_date"`
	CertificatePlan  *string        `json:"certificate_plan"`
	Name             *string        `json:"name"`
	MobileNumber     *string        `json:"mobile_no"`
	Address          *string        `json:"address"`
	Email            *string        `json:"email"`
}

type UserSave struct {
	RegisteredDate  *string `json:"registered_date"`
	CertificatePlan *string `json:"certificate_plan"`
	Name            *string `json:"name"`
	MobileNumber    *string `json:"mobile_no"`
	Address         *string `json:"address"`
	Email           *string `json:"email"`
	CreatedBy       *string `json:"created_by"`
}
