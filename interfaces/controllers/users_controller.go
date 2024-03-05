package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
	usecase "github.com/eduardospek/go-clean-arquiteture/usecases"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(interactor usecase.UserInteractor) *UserController {
	return &UserController{ Interactor: interactor }
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request ) {	
	var user entity.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	err := controller.Interactor.Create(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	w.WriteHeader(http.StatusCreated)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request ) {		
	id := r.URL.Query().Get("id")
	user, err := controller.Interactor.GetById(id)	

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	json.NewEncoder(w).Encode(user)
}