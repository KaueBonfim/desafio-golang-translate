package translate

func (ob *Translate) validationOrder(list_sequence []int){
	last:=0
	for _,num := range list_sequence{
		if last < num && (last == 5 || last == 50 || last == 500){
			panic("Error sintax ")
		}else{
			last=num
		}
	} 
	
}

