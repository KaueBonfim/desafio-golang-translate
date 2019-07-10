package entrypoints

import (
	"fmt"// Para respostas no terminal
)

type Cmd struct {
	Type       string
	Params     map[string]string
	
}

var Help = `
Translate kwego to alphanumeric!!
runner: (./app or app.exe or go run app.go) [COMAND]

Comands: 
	cmd(Default): Translate by alphanumeric command line.

	api: Translate with API: POST localhost:5080/translate

	web: Use a web interface to do kwegonian numbers translation localhost:8080/translate

`

func (cli *Cmd) Run(args []string) {
	
	if len(args) > 0 {
		
		
		cli.Type=args[0]

		// Verifica qual o argumento
		switch cli.Type {
			case "cmd":// linha de comando
				input_cmd()
			case "api":// api na porta 5080
				input_api()
			case "web":// pagina web na porta 8080
				input_web()
			case "help":// help da linha de comando
				fmt.Println(Help)
				return
			default:// comandos invalidos
				fmt.Println("Invavlid Option, set help")
				return 
		}
		

	} else {
		input_cmd()// o padrao e a linha de comando
	}

}

