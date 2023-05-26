package mockdb

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockDB struct {
	DbMock  *sql.DB
	SqlMock sqlmock.Sqlmock
}

func NewMockDB() *MockDB {
	db, sqlMock, err := sqlmock.New()
	if err != nil {
		log.Fatal("failed to create mock database, err: ", err)
	}
	return &MockDB{
		DbMock:  db,
		SqlMock: sqlMock,
	}
}
