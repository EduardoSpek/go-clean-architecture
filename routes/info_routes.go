package routes

import "github.com/eduardospek/go-clean-architecture/interfaces/controllers"

func (r *Router) InfoRouter(infocontroller *controllers.InfoController) {

	r.mux.HandleFunc("/createinfo", infocontroller.CreateInfo).Methods("POST")
	r.mux.HandleFunc("/info/{id}", infocontroller.UpdateInfo).Methods("PATCH")

}