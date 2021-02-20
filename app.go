package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) GetDonors(w http.ResponseWriter, r *http.Request) {
	donors, err := GetDonors(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, donors)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/donors", a.GetDonors).Methods("GET")
	/*a.Router.HandleFunc("/donor", a.CreateDonor).Methods("POST")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.GetDonor).Methods("GET")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.UpdateDonor).Methods("PUT")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.DeleteDonor).Methods("DELETE")*/
}
