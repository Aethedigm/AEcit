package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/api"
	"main/data"
	"main/handlers"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var (
	Api     api.Api
	Handler handlers.Handlers
	Models  data.Model
)

func main() {
	// Create session
	Sess := scs.New()
	Sess.Store = &data.SessionStore{}

	// Create routes
	routes := MakeRoutes(Sess)

	err := godotenv.Load(GetRootPath() + "/.env")
	if err != nil {
		log.Println("Go .Env", err)
		return
	}

	// Connect to DB
	DB, err := ConnectToDB()
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize Api
	Api.Init(DB, Sess)

	// Initialize Handlers
	Handler.Init(GetRootPath(), DB, Sess, &Api)

	// Models init
	Models.New(DB)

	// Print VERSION
	log.Println("VERSION:", os.Getenv("VERSION"))

	if os.Getenv("PRODUCTION") == "true" {
		startRedirectServer()

		log.Println("Listening on port 443 with TLS/SSL")
		http.ListenAndServeTLS(":443", os.Getenv("SECURE_CERT"), os.Getenv("SECURE_KEY"), routes)
	} else {
		port := os.Getenv("PORT")
		log.Println("Listening on port:", port)
		http.ListenAndServe(":"+port, routes)
	}
}

func startRedirectServer() {
	rerouteServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      http.HandlerFunc(redirectToHTTPS),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		rerouteServer.ListenAndServe()
	}()
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, os.Getenv("URL")+r.RequestURI, http.StatusMovedPermanently)
}

func ConnectToDB() (*sql.DB, error) {
	dsn := BuildDSN()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func BuildDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s timezone=UTC connect_timeout=5 password=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASS"),
	)
}

func GetRootPath() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return root
}
