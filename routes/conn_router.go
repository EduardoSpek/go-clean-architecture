package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mux *mux.Router
}

func NewRouter() *Router {
	return &Router{ mux: mux.NewRouter() }
}

func (r *Router) Start(port string) {
	fmt.Println("O Servidor foi iniciado na porta "+ port)
	log.Fatal(http.ListenAndServe(port, r.mux))	
}