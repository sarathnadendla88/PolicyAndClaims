package user

import (
	"fmt"
	"policy/infra/datasources/mysql/db"
	"policy/utils/errors"
)

const (
	insertUser  = "INSERT INTO users (id, created_by, created_datetime, is_active, registered_date,certificate_plan,name,mobile_no,address,email) VALUES(?,?,?,?,?,?,?,?,?,?)"
	getUserById = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, registered_date,certificate_plan,name,mobile_no,address,email from users WHERE id=?"
	getAllUsers = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, registered_date,certificate_plan,name,mobile_no,address,email from users"
)

func (user *User) Save() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := db.Client.Prepare(insertUser)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.Id, user.CreatedBy, user.CreatedDateTime, user.IsActive, user.RegisteredDate, user.CertificatePlan, user.Name, user.MobileNumber, user.Address, user.Email)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save user :%s", err.Error()))
	}
	return nil
}

func (user *User) GetAllUsers() ([]User, *errors.RestErr) {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if user is in DB
	stmt, err := db.Client.Prepare(getAllUsers)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get user rows here
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	// loop user rows
	results := make([]User, 0)
	for rows.Next() {
		//log.Print("new row")
		var user User
		if err := rows.Scan(&user.Id, &user.CreatedBy, &user.CreatedDateTime,
			&user.ModifiedBy, &user.ModifiedDateTime, &user.IsActive,
			&user.RegisteredDate, &user.CertificatePlan, &user.Name, &user.MobileNumber, &user.Address, &user.Email); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequest("No users found.")
	}
	return results, nil
}
func (user *User) GetUserById() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if user is in DB
	stmt, err := db.Client.Prepare(getUserById)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get user rows here
	row := stmt.QueryRow(user.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	if err := row.Scan(&user.Id, &user.CreatedBy, &user.CreatedDateTime,
		&user.ModifiedBy, &user.ModifiedDateTime, &user.IsActive,
		&user.RegisteredDate, &user.CertificatePlan, &user.Name, &user.MobileNumber, &user.Address, &user.Email); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
