package translate

import "strings"// Manipula strings
import "fmt"//Trabalha com a contextos de substituicao de strings

// Struc Para traducao
type Translate struct{
	Input []int `json:"numbers"`
	Text string `json:"input"`
	Output int `json:"output"`
	Invalids []string `json:"invalids"`
}

func (ob *Translate) Getvalues() []int {
	
	block_slipt:=strings.Split(ob.Text," ")// Retira a string e forma um array com os espacos
	for i:=0; i < len(block_slipt); i++{
		//Verifica caracter a caracter se encaixa com os alfanumericos em kwegonian 
		switch block_slipt[i] {
			case "kil":
				ob.Input=append(ob.Input,1)
			case "jin":
				ob.Input=append(ob.Input,5)
			case "pol":
				ob.Input=append(ob.Input,10)
			case "kilow":
				ob.Input=append(ob.Input,50)
			case "jij":
				ob.Input=append(ob.Input,100)
			case "jinjin":
				ob.Input=append(ob.Input,500)
			case "polsx":
				ob.Input=append(ob.Input,1000)
			default:
				ob.get_invalid(block_slipt[i])
		}
		
	}
	//retorna as entradas
	return ob.Input

}

func (ob *Translate) Enumeric(){
	//Faz as comas dos numeros em ordem verificando a logica dos numeros romanos
	if len(ob.Input) > 0{
		ob.validationOrder(ob.Input)
		ultimo:=ob.Input[len(ob.Input)-1]
		for i:=len(ob.Input)-1; i >-1;i-- {
			if ob.Input[i] >= ultimo  {
				ob.Output+=ob.Input[i]
				ultimo=ob.Input[i]
			}else{
				ob.Output-=ob.Input[i]
				ultimo=ob.Input[i]
			}
		}
	}else{
		fmt.Println("No character found")// Se não foi encontrado nenhum alfanumerico entao e retornado a saida nula
	}
}