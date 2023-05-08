package admin

import (
	"fmt"
	"policy/infra/datasources/mysql/db"
	"policy/utils/errors"
)

const (
	insertLocation  = "INSERT INTO locations (id, created_by, created_datetime, is_active, company_name,address,latitude,longitude,description) VALUES(?,?,?,?,?,?,?,?,?)"
	getLocationById = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active,company_name,address,latitude,longitude,description from locations WHERE id=?"
	getAllLocations = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, company_name,address,latitude,longitude,description from locations"
)

func (location *Location) Save() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := db.Client.Prepare(insertLocation)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save location :%s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(location.Id, location.CreatedBy, location.CreatedDateTime, location.IsActive, location.CompanyName, location.Address, location.Latitude, location.Longitude, location.Description)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save location :%s", err.Error()))
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save location :%s", err.Error()))
	}
	return nil
}

func (location *Location) GetAllLocations() ([]Location, *errors.RestErr) {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if location is in DB
	stmt, err := db.Client.Prepare(getAllLocations)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get location rows here
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	// loop through rows
	results := make([]Location, 0)
	for rows.Next() {
		//log.Print("new row")
		var location Location
		if err := rows.Scan(&location.Id, &location.CreatedBy, &location.CreatedDateTime,
			&location.ModifiedBy, &location.ModifiedDateTime, &location.IsActive,
			&location.CompanyName, &location.Address, &location.Latitude, &location.Longitude, &location.Description); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, location)
	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequest("No Location found.")
	}
	return results, nil
}
func (location *Location) GetLocationById() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if location is in DB
	stmt, err := db.Client.Prepare(getLocationById)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get location rows here
	row := stmt.QueryRow(location.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	if err := row.Scan(&location.Id, &location.CreatedBy, &location.CreatedDateTime,
		&location.ModifiedBy, &location.ModifiedDateTime, &location.IsActive,
		&location.CompanyName, &location.Address, &location.Latitude, &location.Longitude, &location.Description); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
