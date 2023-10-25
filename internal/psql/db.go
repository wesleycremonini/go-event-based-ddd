package psql

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type database struct {
	*pgxpool.Pool
}

var db *database

func New(dsn string) (*database, error) {
	newdb, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		zap.L().Error("Unable to connect to database:", zap.String("error", err.Error()))
		os.Exit(1)
	}

	newdb.Config().MaxConnIdleTime = 5 * time.Minute
	newdb.Config().MaxConnLifetime = 2 * time.Hour
	newdb.Config().MaxConns = 25
	newdb.Config().MinConns = 5

	newdb.Config().HealthCheckPeriod = 1 * time.Minute

	db = &database{newdb}
	return &database{newdb}, nil
}
