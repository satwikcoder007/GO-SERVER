package handlers

import (
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to the home page .."))
}
