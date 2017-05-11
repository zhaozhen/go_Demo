package main

import "fmt"

//func test(x,y int,s string)(int ,string){
//	n:=x+y//类型相同的相邻参数可以合并
//	return n,fmt.Sprintf(s,n)//多返回值必须用括号
//}
//func main() {
//	k,v:=test(1,2,"sum:%d")
//	println(k,v)
//}
func test(s string,n ...int) string {
	var x int
	for _,i:=range n{
		x+=i
	}
	return fmt.Sprintf(s,x)
}
func main()  {
	//println(test("sum:%d",1,2,3,4))
	s:=[]int{1,2,3,4}
	println(test("sum:%d",s...))

}