package entrypoints

import (
	"github.com/KaueBonfim/desafio-golang-translate/api" // Trabalha com a entrada da API
)

func input_api(){
	/*
	Inicia o servico da API na porta 5080

	*/
	runner_app := api.Api{}
	runner_app.Init()
	runner_app.Run(":5080")
}