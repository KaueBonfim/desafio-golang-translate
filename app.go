package main

// Importando as entradas modulos de entrada e a lib os 
import (
	cmd "github.com/KaueBonfim/desafio-golang-translate/entrypoints"
	"os"
)

func main() {
	/*
	
	Retirando argumento da linha de comando
	
	Acionando a execução
	*/
	params := os.Args[1:]
	cli := cmd.Cmd{Params:make(map[string]string)}
	cli.Run(params)
}
