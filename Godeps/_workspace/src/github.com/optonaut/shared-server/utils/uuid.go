package utils

import (
	"database/sql/driver"

	"github.com/satori/go.uuid"
)

type UUID string

func (u UUID) String() string {
	return string(u)
}

func GenerateUUID() UUID {
	return UUID(uuid.NewV4().String())
}

func (u *UUID) Scan(value interface{}) error { *u = UUID(value.([]uint8)); return nil }
func (u UUID) Value() (driver.Value, error)  { return string(u), nil }
