package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/donors", a.GetDonors).Methods("GET")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.GetDonor).Methods("GET")
	a.Router.HandleFunc("/donor", a.CreateDonor).Methods("POST")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.UpdateDonor).Methods("PUT")
	a.Router.HandleFunc("/donor/{id:[0-9]+}", a.DeleteDonor).Methods("DELETE")
}
