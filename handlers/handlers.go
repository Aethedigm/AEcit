package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type Handlers struct {
	JetViews *jet.Set
	DB       *sql.DB
	Sessions *scs.SessionManager
}

func (h *Handlers) Render(w http.ResponseWriter, r *http.Request, view string, variables interface{}) error {
	var vars jet.VarMap

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	t, err := h.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		return err
	}

	if err = t.Execute(w, vars, nil); err != nil {
		return err
	}

	return nil
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	count := 0
	if h.Sessions.Exists(r.Context(), "count") {
		count = h.Sessions.GetInt(r.Context(), "count")
		count++
	}

	h.Sessions.Put(r.Context(), "count", count)

	vars := jet.VarMap{}
	vars.Set("count", count)

	err := h.Render(w, r, "home", vars)
	if err != nil {
		log.Println(err)
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}
}
