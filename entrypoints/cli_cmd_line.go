package entrypoints


import (
	"fmt"
	"bufio"
	"os"
	app "github.com/KaueBonfim/desafio-golang-translate/translate" // Lib para traducao
)
	
func input_cmd(){	
	/*
	
	1º Vai verificar a leitura da linha de comando apos a entrada do Input
	2º A guarda o enter e Cria a struct para traducao
	3º Retira os valorese
	4º Faz a traducao
	5º Caso seja maior que 4000 que sao a abase dos numeros romanos volta o erro
	6º Retorna a tradução e a quantidade de palavras que nao sao alfanumericos
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	text, _ := reader.ReadString('\n')
	translate:=app.Translate{Input:[]int{},Text:text[:len(text)-2]}
	translate.Getvalues()
	translate.Enumeric()
	if translate.Output < 4000{
		fmt.Println("Output is =>",translate.Output)
		fmt.Println("No translate your char =>",translate.Invalids)
	}else{
		fmt.Println("This is your number that you have tried to write and is greater than 3999")
	}
	
}
