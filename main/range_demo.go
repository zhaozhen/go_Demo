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
	//rangedo()
	rangedo2()

}
func rangedo2()  {
	s:=[]int{1,2,3,4,5}
	//对复制的对象进行修改
	for i,v:=range s{
		if i==0{
			s=s[:3]
			s[2]=100
		}
		println(i,v)
	}

}
func rangedo() [3]int {
	a:=[3]int{0,1,2}
	//这里证明range对象会复制
	for i,v:= range a{

		if i==0{
			a[1],a[2]=999,999
			fmt.Println(a)
			return a
		}
		a[i]=v+100
	}
	//fmt.Println(a)
	return a
}