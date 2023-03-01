package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/data"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type Controllers struct {
	JetViews *jet.Set
	DB       *sql.DB
	Sessions *scs.SessionManager
	Models   data.Model
}

func (c *Controllers) TestUser(w http.ResponseWriter, r *http.Request) {
	usr, err := c.Models.Users.GetByEmail("aethedigm@gmail.com")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (c *Controllers) Render(w http.ResponseWriter, r *http.Request, view string, variables interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	t, err := c.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		return err
	}

	if err = t.Execute(w, vars, nil); err != nil {
		return err
	}

	return nil
}

func (c *Controllers) Home(w http.ResponseWriter, r *http.Request) {
	count := 0
	if c.Sessions.Exists(r.Context(), "count") {
		count = c.Sessions.GetInt(r.Context(), "count")
		count++
	}

	c.Sessions.Put(r.Context(), "count", count)

	vars := jet.VarMap{}
	vars.Set("count", count)

	err := c.Render(w, r, "home", vars)
	if err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}
