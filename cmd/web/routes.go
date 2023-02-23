package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	//Create a multiplexer
	router := httprouter.New()
	//Create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //Changed this around to implement the rest capability
	router.HandlerFunc(http.MethodGet, "/create", app.Greeting)
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.HandlerFunc(http.MethodPost, "/create", app.MessageCreate)

	return router
}
