package service

import (
	"database/sql"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

type Service interface {
	db.Store
}

type service struct {
	store   db.Store
	queries *db.Queries
	sqlDB   *sql.DB
}

func New(sqlDB *sql.DB) Service {
	return &service{
		store:   db.NewStore(sqlDB),
		queries: db.New(sqlDB),
		sqlDB:   sqlDB,
	}
}
