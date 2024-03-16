package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
	usecase "github.com/eduardospek/go-clean-arquiteture/usecases"
)

type InfoController struct {
	InfoInteractor usecase.InfoInteractor
}

func NewInfoController(infointeractor usecase.InfoInteractor) *InfoController {
	return &InfoController{InfoInteractor: infointeractor}
}

func (controller *InfoController) CreateInfo(w http.ResponseWriter, r *http.Request) {
	var info entity.InfoDTO

	_ = json.NewDecoder(r.Body).Decode(&info)
	newinfo, err := controller.InfoInteractor.CreateInfo(info.Id_user, info.Cabelo, info.Olhos, info.Pele, info.Corpo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newinfo)
}