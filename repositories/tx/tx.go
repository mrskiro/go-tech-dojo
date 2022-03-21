package tx

import (
	"context"
	"database/sql"
)

type Repository interface {
	Start(ctx context.Context, fn func(ctx context.Context) error) error
}

type repository struct {
	db *sql.DB
}

func NewTxRepository(db *sql.DB) Repository {
	return repository{
		db: db,
	}
}

func (r repository) Start(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	err = fn(ctx)
	if err != nil {
		err := tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
