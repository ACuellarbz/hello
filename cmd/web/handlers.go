package main

import (
	"net/http"

	"github.com/ACuellarbz/hello/helpers"
)

// Create our handler projects
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {

}

// Create our handler functions
func (app *application) About(w http.ResponseWriter, r *http.Request) {

}

// Create our handler functions
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//see a web page displayed, to do this i need to render a template.

	helpers.RenderTemplates(w, "./static/html/home.page.tmpl")

}

// Create our handler functions
func (app *application) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
