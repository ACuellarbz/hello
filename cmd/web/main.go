// route = /greeting  handler = greeting
// message will be " Welcome to my page"

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ACuellarbz/hello/internal/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// create a new type
type application struct {
	question models.QuestionModel //application type containts a "question" of type Models.QuestionModel
}

// MAIN PROGRAM BELOW
func main() {
	//create a flag for speifing the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String(" dsn", os.Getenv("BUSSES_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()
	//create an instance of a connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	//create an instance of the application type
	app := &application{
		question: models.QuestionModel{DB: db},
		
	}

	defer db.Close()
	log.Println("Database connection pool established")
	//create our own server (replacing listenandserve()
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	//use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //specify timeout
	defer cancel()                                                          //+Defer attached to a function and executes as the last thing
	//lets ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
