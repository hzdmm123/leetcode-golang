package kit

func StringToCharArray(str string) (strs []string){
	for _,element := range str {
		strs = append(strs,string(element))
	}
	return
}
