package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
	usecase "github.com/eduardospek/go-clean-arquiteture/usecases"
	"github.com/gorilla/mux"
)

type UserController struct {
	UserInteractor usecase.UserInteractor
}

func NewUserController(userinteractor usecase.UserInteractor) *UserController {
	return &UserController{ UserInteractor: userinteractor }
}

func (controller *UserController) CreateUser(w http.ResponseWriter, r *http.Request ) {	
	var user entity.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	newuser, err := controller.UserInteractor.CreateNewUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	json.NewEncoder(w).Encode(newuser)
}
func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request ) {	
	var user entity.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	UpdateUser, err := controller.UserInteractor.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	json.NewEncoder(w).Encode(UpdateUser)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	user, err := controller.UserInteractor.GetById(id)	

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	json.NewEncoder(w).Encode(user)
}

func (controller *UserController) UserList(w http.ResponseWriter, r *http.Request ) {			
	users, err := controller.UserInteractor.UserList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	json.NewEncoder(w).Encode(users)
}

func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request ) {			
	vars := mux.Vars(r)
	id := vars["id"]
	
	err := controller.UserInteractor.DeleteUser(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusOK)
}