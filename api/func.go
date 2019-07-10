package api

import (
	"net/http"
	"encoding/json"
	app "github.com/KaueBonfim/desafio-golang-translate/translate"
)

func (api *Api) Translater(w http.ResponseWriter, r *http.Request) {
	var translate app.Translate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&translate); err != nil {
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	translate.Getvalues()
	translate.Enumeric()
	if translate.Text == ""{
		http.Error(w, "Invalid Input", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(translate)
}

func swagger(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    http.ServeFile(w, r, "swagger.json")
}

