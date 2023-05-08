package admin

import (
	"fmt"
	"policy/infra/datasources/mysql/db"
	"policy/utils/errors"
)

const (
	insertUser  = "INSERT INTO agent (id, created_by, created_datetime, is_active, name,mobile_no,email) VALUES(?,?,?,?,?,?,?)"
	getUserById = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active,name,mobile_no,email from agent WHERE id=?"
	getAllUsers = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, name,mobile_no,email from agent"
)

func (admin *UserData) Save() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := db.Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(admin.Id, admin.CreatedBy, admin.CreatedDateTime, admin.IsActive, admin.Name, admin.MobileNo, admin.Email)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	return nil
}

func (user *UserData) GetAllUsers() ([]UserData, *errors.RestErr) {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if admin is in DB
	stmt, err := db.Client.Prepare(getAllUsers)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get admin rows here
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	// loop through rows
	results := make([]UserData, 0)
	for rows.Next() {
		//log.Print("new row")
		var admin UserData
		if err := rows.Scan(&admin.Id, &admin.CreatedBy, &admin.CreatedDateTime,
			&admin.ModifiedBy, &admin.ModifiedDateTime, &admin.IsActive,
			&admin.Name, &admin.MobileNo, &admin.Email); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, admin)
	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequest("No User found.")
	}
	return results, nil
}
func (admin *UserData) GetUserById() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if temp is in DB
	stmt, err := db.Client.Prepare(getUserById)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get admin rows here
	row := stmt.QueryRow(admin.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	if err := row.Scan(&admin.Id, &admin.CreatedBy, &admin.CreatedDateTime,
		&admin.ModifiedBy, &admin.ModifiedDateTime, &admin.IsActive,
		&admin.Name, &admin.MobileNo, &admin.Email); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
