package main

func main() {
	x := "abc"
	for i, n := 0, len(x); i < n; i++ {
		println(i, x[i])
	}
	//n:=len(x)
	//for n>0{
	//	n--
	//	println(x[n])
	//
	//}
	//for {
	//	println(x)
	//}
}
func length(s string) int {
	println("call length-------------------")
	return len(s)
}
