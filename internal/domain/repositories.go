package domain

import (
	"context"

	"github.com/google/uuid"
)

type CustomerRepository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*Customer, error)
	GetByID(ctx context.Context, id int) (*Customer, error)
	GetByToken(ctx context.Context, token string) (*Customer, error)
}

type CompanyRepository interface {
	GetByCustomerID(ctx context.Context, customerID int) (*Company, error)
	GetByID(ctx context.Context, id int) (*Company, error)
}

type UserRepository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*User, error)
}