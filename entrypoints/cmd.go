package entrypoints

import (
	"fmt"
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


		switch cli.Type {
			case "cmd":
				input_cmd()
			case "api":
				input_api()
			case "web":
				input_web()
			case "help":
				fmt.Println(Help)
				return
			default:
				fmt.Println("Invavlid Option")
				return 
		}
		

	} else {
		input_cmd()
	}

}

