package app

import (
	"github.com/gorilla/mux"
	mw "podbilling/users/pkg/middleware"
	"log"
	"net/http"
)

func (app *App) WithAuthMW(h mw.PrivatePageHandler) http.HandlerFunc {
	return mw.JwtAuthMW(app.JwtSignKey, h)
}

func (app *App) Route() {

	r := mux.NewRouter().PathPrefix("/user").Subrouter()
	// r.Schemes("https")

	r.HandleFunc("/me/", app.WithAuthMW(app.Handler.GetSelfUser)).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}/", app.WithAuthMW(app.Handler.GetUserByID)).Methods("GET")
	r.HandleFunc("/workers/", app.WithAuthMW(app.Handler.GetAllWorkers)).Methods("GET")
	r.HandleFunc("/", app.WithAuthMW(app.Handler.CreateUser)).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}/", app.WithAuthMW(app.Handler.UpdateUser)).Methods("PUT")
	r.HandleFunc("/{id:[0-9]+}/", app.WithAuthMW(app.Handler.DeleteUser)).Methods("DELETE")

	(*app).Server.Handler = r

	log.Fatal(app.Server.ListenAndServe())
}
