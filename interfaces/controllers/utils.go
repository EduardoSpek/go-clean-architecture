package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data any, statusCode int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {	
		return errors.New("responseJson: não foi possível converter para json")
	}	

	// Escrevendo a resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonData)

	return nil

}