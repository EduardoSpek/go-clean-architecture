package routes

import (
	"fmt"
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
	r.mux.HandleFunc("/userlist", usercontroller.UserList).Methods("GET")
	r.mux.HandleFunc("/user/{id}", usercontroller.UpdateUser).Methods("PUT")
	r.mux.HandleFunc("/user/{id}", usercontroller.GetUser).Methods("GET")
	r.mux.HandleFunc("/user/{id}", usercontroller.DeleteUser).Methods("DELETE")
	
}

func (r *Router) Start(port string) {
	fmt.Println("O Servidor foi iniciado na porta "+ port)
	log.Fatal(http.ListenAndServe(port, r.mux))	
}