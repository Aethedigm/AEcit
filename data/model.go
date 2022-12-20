package data

import (
	"database/sql"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	_ "github.com/upper/db/v4/adapter/postgresql"
)

type Model struct {
	Users User
}

var db *sql.DB
var upper db2.Session

func (m *Model) New(database *sql.DB) Model {
	var err error
	db = database

	upper, err = postgresql.New(db)
	if err != nil {
		panic(err)
	}

	return Model{
		Users: User{},
	}
}
