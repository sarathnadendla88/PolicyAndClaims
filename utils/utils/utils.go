package utils

import (
	"policy/utils/errors"
	"database/sql/driver"
	errors2 "errors"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

func GetRandomString() (string, *errors.RestErr) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", errors.NewInternalServerError(fmt.Sprintf("Error in creating id: %s", err.Error()))
	}
	return id.String(), nil
}

func GetCreatedBy() *string {
	var admin = "admin"
	return &admin
}

type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b BitBool) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors2.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}

func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
