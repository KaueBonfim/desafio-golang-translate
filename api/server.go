package api

import (
	"net/http"// Utilizado para levantar o servidor
	"github.com/gorilla/mux" // Trabalhar as rotas da api
	"log"// Gerar o log do desligamento do servidor
	"fmt"// Para respostas no terminal
)

// Modelo de objeto que controla as rotas
type Api struct {
	router   *mux.Router
}


func (app *Api) Init() {
	/*
	metodo para tradução:
	1º Cria o modelo de rotas com Gorilla mux
	3º Construir o caminho translate para POST no controller Translate
	*/
	app.router = mux.NewRouter()
	app.router.HandleFunc("/translate", app.Translater).Methods("POST")
}

func (app *Api) Run(addr string) {
	/*
	Execucao do servidor
	*/
	fmt.Println("Starting API on :5080/translate ")
	log.Fatal(http.ListenAndServe(addr, app.router))
}


