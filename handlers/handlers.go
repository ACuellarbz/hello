package handlers

import (
	"fmt"
	"net/http"
	"time"
)

//Create our handler projects
func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my page"))
}

//Create our handler functions
func About(w http.ResponseWriter, r *http.Request) {
	day := time.Now().Weekday()
	w.Write([]byte(fmt.Sprintf("Have a good %s", day)))
}

//Create our handler functions
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to my home page"))
}

//Create our handler functions
func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("message created..."))
}
