package controllers

import (
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
		http.Error(w, "Usuário não encontrado", http.StatusInternalServerError)
		return
	}

	ResponseJson(w, userinfo, http.StatusOK)	
}