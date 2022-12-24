package calc

func Pow(a ,b int) int {
	res:=1
	for i:=0; i<b;i++{
		res=res*a
	}


	return res
}
