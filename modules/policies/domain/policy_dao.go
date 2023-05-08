package admin

import (
	"fmt"
	"policy/infra/datasources/mysql/db"
	"policy/utils/errors"
)

const (
	insertPolicy   = "INSERT INTO policies (id, created_by, created_datetime, is_active, policy_name,amount,duration_of_policy,final_reedemable_amount,description) VALUES(?,?,?,?,?,?,?,?,?)"
	getPolicyById  = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active,policy_name,amount,duration_of_policy,final_reedemable_amount,description from policies WHERE id=?"
	getAllPolicies = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active,policy_name,amount,duration_of_policy,final_reedemable_amount,description from policies"
)

func (policy *Policy) Save() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := db.Client.Prepare(insertPolicy)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save policy :%s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(policy.Id, policy.CreatedBy, policy.CreatedDateTime, policy.IsActive, policy.PolicyName, policy.Amount, policy.DurationOfPolicy, policy.FinalReedemableAmount, policy.Description)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save policy :%s", err.Error()))
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save policy :%s", err.Error()))
	}
	return nil
}

func (policy *Policy) GetAllPolicies() ([]Policy, *errors.RestErr) {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if policy is in DB
	stmt, err := db.Client.Prepare(getAllPolicies)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get policy rows here
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	// loop policy rows
	results := make([]Policy, 0)
	for rows.Next() {
		//log.Print("new row")
		var policy Policy
		if err := rows.Scan(&policy.Id, &policy.CreatedBy, &policy.CreatedDateTime,
			&policy.ModifiedBy, &policy.ModifiedDateTime, &policy.IsActive,
			&policy.PolicyName, &policy.Amount, &policy.DurationOfPolicy, &policy.FinalReedemableAmount, &policy.Description); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, policy)
	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequest("No policies found.")
	}
	return results, nil
}
func (policy *Policy) GetPolicyById() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if policy is in DB
	stmt, err := db.Client.Prepare(getPolicyById)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get policy rows here
	row := stmt.QueryRow(policy.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	if err := row.Scan(&policy.Id, &policy.CreatedBy, &policy.CreatedDateTime,
		&policy.ModifiedBy, &policy.ModifiedDateTime, &policy.IsActive,
		&policy.PolicyName, &policy.Amount, &policy.DurationOfPolicy, &policy.FinalReedemableAmount, &policy.Description); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
