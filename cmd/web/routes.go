package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//Create a multiplexer
	mux := httprouter.New()
	//Create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //Changed this around to implement the rest capability
	mux.HandlerFunc(http.MethodGet, "/create", app.Greeting)
	mux.HandlerFunc(http.MethodGet, "/about", app.About)
	mux.HandlerFunc(http.MethodGet, "/", app.Home)
	mux.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)

	return mux
}
