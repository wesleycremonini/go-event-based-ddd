package psql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/wesleycremonini/go-event-based-ddd/internal/domain"
)

type CustomerRepository struct {
	*database
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.Customer, error) {
	rows, err := r.Query(ctx, `SELECT * FROM customers WHERE uuid = $1 LIMIT 1`, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	client, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *CustomerRepository) GetByID(ctx context.Context, id int) (*domain.Customer, error) {
	rows, err := r.Query(ctx, `SELECT * FROM customers WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	client, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *CustomerRepository) GetByToken(ctx context.Context, token string) (*domain.Customer, error) {
	rows, err := r.Query(ctx, `SELECT * FROM customers WHERE token = $1 LIMIT 1`, token)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	client, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Customer])
	if err != nil {
		return nil, err
	}

	return &client, nil
}
