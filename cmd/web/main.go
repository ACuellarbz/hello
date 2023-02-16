// route = /greeting  handler = greeting
// message will be " Welcome to my page"

package main

import (
	"flag"
	"log"
	"net/http"
)

// create a new type
type application struct {
}

// MAIN PROGRAM BELOW
func main() {

	//create a flag for speifing the port number when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	flag.Parse()

	//create an instance of the application type
	app := &application{}

	//create our own server (replacing listenandserve()
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
