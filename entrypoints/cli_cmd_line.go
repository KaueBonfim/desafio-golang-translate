package entrypoints


import (
	"fmt"
	"bufio"
	"os"
	app "github.com/KaueBonfim/desafio-golang-translate/translate"
)
	
func input_cmd(){	
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
