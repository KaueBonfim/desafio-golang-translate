package entrypoints

import (
	"github.com/KaueBonfim/desafio-golang-translate/api"
)

func input_api(){
	runner_app := api.Api{}
	runner_app.Init()
	runner_app.Run(":5080")
}