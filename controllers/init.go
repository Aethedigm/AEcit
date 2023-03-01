package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

func (c *Controllers) Init(rootPath string, db *sql.DB, Sess *scs.SessionManager) {
	c.JetViews = jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	c.DB = db

	c.Sessions = Sess

	c.Sessions.Lifetime = 24 * time.Hour
	c.Sessions.IdleTimeout = 120 * time.Minute
	c.Sessions.Cookie.Name = os.Getenv("COOKIE_NAME")
	c.Sessions.Cookie.Domain = os.Getenv("COOKIE_DOMAIN")
	c.Sessions.Cookie.HttpOnly = false
	c.Sessions.Cookie.Persist = os.Getenv("COOKIE_PERSIST") == "true"
	c.Sessions.Cookie.SameSite = http.SameSiteStrictMode
	c.Sessions.Cookie.Secure = os.Getenv("COOKIE_SECURE") == "true"
}
