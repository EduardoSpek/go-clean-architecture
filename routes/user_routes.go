package routes

import (
	"log"
	"net/http"

	"github.com/eduardospek/go-clean-arquiteture/interfaces/controllers"
	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	return &Router{ mux: mux.NewRouter() }
}

func (r *Router) SetupRouter( usercontroller *controllers.UserController) {
	r.mux.HandleFunc("/createuser", usercontroller.CreateUser).Methods("POST")
	r.mux.HandleFunc("/user", usercontroller.GetUser).Methods("GET")
	r.mux.HandleFunc("/userlist", usercontroller.UserList).Methods("GET")
}

func (r *Router) Start(port string) {
	log.Fatal(http.ListenAndServe(port, r.mux))
}