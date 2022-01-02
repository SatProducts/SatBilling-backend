package app

import (
	"github.com/gorilla/mux"
	"log"
)

func (app *App) Route() {

	r := mux.NewRouter().PathPrefix("/auth").Subrouter()
	// r.Schemes("https")

	r.HandleFunc("/login/", app.Handler.Login).Methods("POST")

	(*app).Server.Handler = r

	log.Fatal(app.Server.ListenAndServe())
}
