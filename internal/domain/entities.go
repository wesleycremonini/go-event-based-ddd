package domain

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Company struct {
	ID         int       `db:"id"`
	UUID       uuid.UUID `db:"uuid"`
	CNPJ       string    `db:"cnpj"`
	CustomerID int       `db:"customer_id"`
}

type User struct {
	ID         int       `db:"id"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	UUID       uuid.UUID `db:"uuid"`
	CustomerID int       `db:"customer_id"`
}

type Customer struct {
	ID    int         `db:"id"`
	UUID  uuid.UUID   `db:"uuid"`
	Token null.String `db:"token"`
}