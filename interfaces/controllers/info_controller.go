package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
)

type InfoController struct {
	InfoInteractor usecase.InfoInteractor
}

func NewInfoController(infointeractor usecase.InfoInteractor) *InfoController {
	return &InfoController{InfoInteractor: infointeractor}
}

func (controller *InfoController) CreateInfo(w http.ResponseWriter, r *http.Request) {
	var info entity.InfoInput	
	
	_ = json.NewDecoder(r.Body).Decode(&info)

	newinfo, err := controller.InfoInteractor.CreateInfo(info)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newinfo)
}