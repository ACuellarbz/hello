// route = /greeting  handler = greeting
// message will be " Welcome to my page"

package main

import (
	"log"
	"net/http"

	"github.com/ACuellarbz/hello/handlers"
)

// MAIN PROGRAM BELOW
func main() {
	mux := http.NewServeMux()
	//Cresate a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))
	

	mux.HandleFunc("/greeting", handlers.Greeting)
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/message/create", handlers.MessageCreate)
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
