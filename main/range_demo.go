package main

import "fmt"

func main() {
	//s:="abc"
	//for i:=range  s{
	//	println(s[i])
	//}
	//
	//for _,c:=range s{
	//	println(c)
	//}
	//
	//for range s{
	//	//...
	//}
	//m:=map[string]int{"a":1,"b":2}
	//
	//for k,v:=range m{
	//	println(k,v)
	//}
	a:=[3]int{0,1,2}
	for i,v:= range a{
		if i==0{
			a[1],a[2]=999,999
			fmt.Println(a)
		}
		a[i]=v+100
	}

	fmt.Println(a)
}