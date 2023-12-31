package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leonardfreitas/go-client-server-api/server/services"
)

func GetDollarControler(w http.ResponseWriter, r *http.Request) {
	result, err := services.GetDollarService()

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	jsonPayload, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	w.Write(jsonPayload)
}
