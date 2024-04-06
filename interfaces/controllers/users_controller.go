package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
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
	
	ResponseJson(w, newuser, http.StatusCreated)
}

func (controller *UserController) UpdateUser(w http.ResponseWriter, r *http.Request ) {	
	vars := mux.Vars(r)
	id := vars["id"]

	var user entity.User

	user.ID = id

	_ = json.NewDecoder(r.Body).Decode(&user)
	UpdateUser, err := controller.UserInteractor.UpdateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	
	ResponseJson(w, UpdateUser, http.StatusOK)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	user, err := controller.UserInteractor.GetById(id)	

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	ResponseJson(w, user, http.StatusOK)
}

func (controller *UserController) UserList(w http.ResponseWriter, r *http.Request ) {			
	users, err := controller.UserInteractor.UserList()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}
	
	ResponseJson(w, users, http.StatusOK)
}

func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request ) {			
	vars := mux.Vars(r)
	id := vars["id"]

	err := controller.UserInteractor.DeleteUser(id)

	msg := map[string]any{ 
		"message": "usuário não encontrado",
		"error": true,
	}

	if err != nil {
		ResponseJson(w, msg, http.StatusOK)
		return
	}

	msg = map[string]any{ 
		"message": "usuário deletado com sucesso!",
		"error": false,
	}

	ResponseJson(w, msg, http.StatusOK)	

}