package psql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/wesleycremonini/go-event-based-ddd/internal/domain"
)

type CompanyRepository struct {
	*database
}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) GetByCustomerID(ctx context.Context, customerID int) (*domain.Company, error) {
	rows, err := r.Query(ctx, `SELECT * FROM companies WHERE customer_id = $1 LIMIT 1`, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	company, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Company])
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (r *CompanyRepository) GetByID(ctx context.Context, id int) (*domain.Company, error) {
	rows, err := r.Query(ctx, `SELECT * FROM companies WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	company, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.Company])
	if err != nil {
		return nil, err
	}

	return &company, nil
}
