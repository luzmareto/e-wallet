package service

import (
	"database/sql"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

type Service interface {
	db.Querier
}

type service struct {
	queries *db.Queries
	sqlDB   *sql.DB
}

func New(sqlDB *sql.DB) Service {
	return &service{
		queries: db.New(sqlDB),
		sqlDB:   sqlDB,
	}
}
