package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/gorilla/mux"
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

	ResponseJson(w, newinfo, http.StatusOK)
}

func (controller *InfoController) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var info entity.InfoInput

	info.ID = id
	
	_ = json.NewDecoder(r.Body).Decode(&info)

	newinfo, err := controller.InfoInteractor.UpdateInfo(info)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJson(w, newinfo, http.StatusOK)

}