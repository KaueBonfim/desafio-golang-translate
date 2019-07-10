package entrypoints

import (
	"github.com/KaueBonfim/desafio-golang-translate/web" // Trabalha com a entrada da WEB
)

func input_web(){
	/*
	Inicia o servico da API na porta 8080
	*/
	web.Run()
}
