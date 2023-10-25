package psql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/wesleycremonini/go-event-based-ddd/internal/domain"
)

type UserRepository struct {
	*database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*domain.User, error) {
	rows, err := r.Query(ctx, `SELECT * FROM users WHERE uuid = $1 LIMIT 1`, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[domain.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}
