package translate


func (ob *Translate) get_invalid(value string){
	ob.Invalids=append(ob.Invalids,value)
}
