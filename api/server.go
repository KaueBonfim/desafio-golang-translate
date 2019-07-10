package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
)

type Api struct {
	router   *mux.Router
}


func (app *Api) Init() {
	app.router = mux.NewRouter()
	app.router.HandleFunc("/translate", app.Translater).Methods("POST")
}

func (app *Api) Run(addr string) {
	fmt.Println("Starting API on :5080/translate ")
	log.Fatal(http.ListenAndServe(addr, app.router))
}


