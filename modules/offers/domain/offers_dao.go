package offers

import (
	"fmt"
	"policy/infra/datasources/mysql/db"
	"policy/utils/errors"
)

const (
	insertOffer  = "INSERT INTO offers (id, created_by, created_datetime, is_active, name,offer_code,description,image_url) VALUES(?,?,?,?,?,?,?,?)"
	getOfferById = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, name,offer_code,description,image_url from offers WHERE id=?"
	getAllOffers = "SELECT id, created_by, created_datetime,modified_by, modified_datetime, is_active, name,offer_code,description,image_url from offers"
)

func (offer *Offers) Save() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := db.Client.Prepare(insertOffer)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save offer :%s", err.Error()))
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(offer.Id, offer.CreatedBy, offer.CreatedDateTime, offer.IsActive, offer.Name, offer.OfferCode, offer.Description, offer.ImageUrl)

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save offer :%s", err.Error()))
	}
	_, err = insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Not able to save offer :%s", err.Error()))
	}
	return nil
}

func (offer *Offers) GetAllOffers() ([]Offers, *errors.RestErr) {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if offer is in DB
	stmt, err := db.Client.Prepare(getAllOffers)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get offer rows here
	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	// loop offer rows
	results := make([]Offers, 0)
	for rows.Next() {
		//log.Print("new row")
		var offer Offers
		if err := rows.Scan(&offer.Id, &offer.CreatedBy, &offer.CreatedDateTime,
			&offer.ModifiedBy, &offer.ModifiedDateTime, &offer.IsActive,
			&offer.Name, &offer.OfferCode, &offer.Description, &offer.ImageUrl); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, offer)
	}
	rows.Close()

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequest("No users found.")
	}
	return results, nil
}
func (offer *Offers) GetOfferById() *errors.RestErr {
	if err := db.Client.Ping(); err != nil {
		panic(err)
	}
	// check if offer is in DB
	stmt, err := db.Client.Prepare(getOfferById)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Query error %s", err.Error()))
	}
	defer stmt.Close()

	// get offer rows here
	row := stmt.QueryRow(offer.Id)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("%s", err.Error()))
	}

	if err := row.Scan(&offer.Id, &offer.CreatedBy, &offer.CreatedDateTime,
		&offer.ModifiedBy, &offer.ModifiedDateTime, &offer.IsActive,
		&offer.Name, &offer.OfferCode, &offer.Description, &offer.ImageUrl); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
