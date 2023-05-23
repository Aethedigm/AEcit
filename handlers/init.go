package handlers

import (
	"database/sql"
	"fmt"
	"main/api"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

func (h *Handlers) Init(rootPath string, db *sql.DB, Sess *scs.SessionManager, api *api.Api) {
	h.JetViews = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	h.DB = db

	h.Sessions = Sess

	h.Sessions.Lifetime = 24 * time.Hour
	h.Sessions.IdleTimeout = 120 * time.Minute
	h.Sessions.Cookie.Name = os.Getenv("COOKIE_NAME")
	h.Sessions.Cookie.Domain = os.Getenv("COOKIE_DOMAIN")
	h.Sessions.Cookie.HttpOnly = false
	h.Sessions.Cookie.Persist = os.Getenv("COOKIE_PERSIST") == "true"
	h.Sessions.Cookie.SameSite = http.SameSiteStrictMode
	h.Sessions.Cookie.Secure = os.Getenv("COOKIE_SECURE") == "true"

	h.Api = *api
}
