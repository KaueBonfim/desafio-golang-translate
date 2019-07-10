package translate

// Retira os valores invalidos e guarda no arrey da struc Translate
func (ob *Translate) get_invalid(value string){
	ob.Invalids=append(ob.Invalids,value)
}
