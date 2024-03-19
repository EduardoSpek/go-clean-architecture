package routes

import "github.com/eduardospek/go-clean-architecture/interfaces/controllers"

func (r *Router) UserInfoRouter(userinfocontroller *controllers.UserInfoController) {

	r.mux.HandleFunc("/userinfo/{id}", userinfocontroller.Get).Methods("GET")

}