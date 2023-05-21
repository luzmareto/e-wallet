package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

type sqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &sqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *sqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("err : %v, rb err : %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
