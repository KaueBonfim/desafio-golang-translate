package web

import (
	"fmt"
	"html/template"
	"net/http"
	app "github.com/KaueBonfim/desafio-golang-translate/translate"
	"strings"
)


func translate(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("./web/template/translate.html")

	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}
	
	r.ParseForm()
	translate:=app.Translate{Input:[]int{},Text:strings.Join(r.Form["input"],"")}
	translate.Getvalues()
	translate.Enumeric()
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, translate)
}

func Run() {
	http.HandleFunc("/translate", translate)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}