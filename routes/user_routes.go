package routes

import (
	"github.com/eduardospek/go-clean-architecture/interfaces/controllers"
)

func (r *Router) UserRouter( usercontroller *controllers.UserController) {
	r.mux.HandleFunc("/createuser", usercontroller.CreateUser).Methods("POST")
	r.mux.HandleFunc("/userlist", usercontroller.UserList).Methods("GET")
	r.mux.HandleFunc("/user/{id}", usercontroller.UpdateUser).Methods("PUT")
	r.mux.HandleFunc("/user/{id}", usercontroller.GetUser).Methods("GET")
	r.mux.HandleFunc("/user/{id}", usercontroller.DeleteUser).Methods("DELETE")	
	
}