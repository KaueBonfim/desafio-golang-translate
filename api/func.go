package api

import (
	"net/http"// Utilizado para request e response
	"encoding/json" // Trabalha traduzindo as entradas e as saidas para a struct de traducao
	app "github.com/KaueBonfim/desafio-golang-translate/translate" //app para traduzir numeros do idioma kwegonian
)

func (api *Api) Translater(w http.ResponseWriter, r *http.Request) {
	/*
	metodo para tradução:
	1º Cria um objeto app para tradução
	2º Recebe um post com o valor Input como texto
	3º Verifica se foi recebido o valor do Input com post
	4º Faz a tradução
	5º Verifica se ouve algum valor traduzido
	6º Retorna a resposta completa da tradução
	*/
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


