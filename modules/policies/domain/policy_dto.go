package admin

import (
	"policy/utils/utils"
)

type Policy struct {
	Id                    string         `json:"id" `
	IsActive              *utils.BitBool `json:"is_active"`
	CreatedDateTime       *string        `json:"created_date_time"`
	ModifiedDateTime      *string        `json:"modified_date_time"`
	ModifiedBy            *string        `json:"modified_by"`
	CreatedBy             *string        `json:"created_by"`
	PolicyName            *string        `json:"policy_name"`
	Amount                *string        `json:"amount"`
	DurationOfPolicy      *string        `json:"duration_of_policy"`
	FinalReedemableAmount *string        `json:"final_reedemable_amount"`
	Description           *string        `json:"description"`
}

type PolicySave struct {
	PolicyName            *string `json:"policy_name"`
	Amount                *string `json:"amount"`
	DurationOfPolicy      *string `json:"duration_of_policy"`
	FinalReedemableAmount *string `json:"final_reedemable_amount"`
	Description           *string `json:"description"`
	CreatedBy             *string `json:"created_by"`
}
