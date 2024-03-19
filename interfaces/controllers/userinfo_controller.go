package controllers

import (
	"encoding/json"
	"net/http"

	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/gorilla/mux"
)

type UserInfoController struct {
	UserInfoInteractor usecase.UserInfoInteractor
}

func NewUserInfoController(userinfointeractor usecase.UserInfoInteractor) *UserInfoController {
	return &UserInfoController{ UserInfoInteractor: userinfointeractor}
}

func (c *UserInfoController) Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]	

	userinfo, err := c.UserInfoInteractor.Get(id)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(userinfo)
	if err != nil {
		http.Error(w, "Erro ao serializar para JSON", http.StatusInternalServerError)
		return
	}

	// Escrevendo a resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}