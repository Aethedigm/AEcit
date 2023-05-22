package api

import (
	"database/sql"
	"main/data"

	"github.com/alexedwards/scs/v2"
)

type Api struct {
	DB       *sql.DB
	Sessions *scs.SessionManager
	Models   data.Model
}

func (a *Api) Init(db *sql.DB, sess *scs.SessionManager) {
	a.DB = db

	a.Sessions = sess
}