package web


import (
	"fmt" // Para respostas no terminal
	"html/template" //Para renderizar template 
	"net/http" //Utilizado para levantar o servidor
	app "github.com/KaueBonfim/desafio-golang-translate/translate" //app para traduzir numeros do idioma kwegonian
	"strings" //manipulacao de strings
)


func translate(w http.ResponseWriter, r *http.Request) {
	/*
	metodo para tradução:
	1º Leitura do template e verificar se a chamada e GET ou POST
	2º Se GET aparecera a modelo de input
	3º Se POST do formulario, então aparece o resultado da tradução
	*/
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
	/*
	Execucao do servidor
	*/

	http.HandleFunc("/translate", translate)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}