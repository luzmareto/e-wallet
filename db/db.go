package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func Connect(config utils.Config) *sql.DB {
	dbConnection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	if err := dbConnection.Ping(); err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	dbConnection.SetMaxOpenConns(500)
	dbConnection.SetMaxIdleConns(50)
	dbConnection.SetConnMaxIdleTime(time.Minute * 10)
	dbConnection.SetConnMaxLifetime(time.Hour * 1)
	return dbConnection
}
