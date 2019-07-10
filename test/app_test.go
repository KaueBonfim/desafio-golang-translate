package test

import (
	"testing"// lib de test
	app "github.com/KaueBonfim/desafio-golang-translate/translate" // Lib para traducao
	"fmt"// Trabalha com a linha de comando
)

func TestTranslate(t *testing.T) {
	//Verifica se a traducao esta correta
	translate:= app.Translate{Input:[]int{},Text:"polsx polsx pol jin kil"}
	translate.Getvalues()
	translate.Enumeric()
	if translate.Output != 2016{
		t.Errorf("The value of translation is wrong the result that was generated from the translation is: %d, when should be 2016",translate.Output)
	}
}


func TestInputs(t *testing.T) {
	//Verifica se a conversao das palavras para numeros

	translate:= app.Translate{Input:[]int{},Text:"polsx polsx pol jin kil"}
	translate.Getvalues()
	value_input:=[]int{1000,1000,10,5,1}
	for i, v := range translate.Input {
        if v != value_input[i] {
			t.Errorf("Your Input is different from the data you returned")
        }
    }
	
}


func TestGetInvalidsChars(t *testing.T) {
	//Verifica os caracteres invalidos estao corretos

	translate:= app.Translate{Input:[]int{},Text:"Testado polsx Com pol Sucesso kil"}
	translate.Getvalues()
	if len(translate.Invalids) == 3{
		for _,v:= range translate.Invalids{
			switch v{
				case "Testado":
					fmt.Print("Testado ")
				case "Com":
					fmt.Print("Com ")
				case "Sucesso":
					fmt.Println("Sucesso")
				default:
					t.Errorf("Non-alphanumeric values were not correctly identified")
			}
		}
	}
}