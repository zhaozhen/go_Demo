package main

//func test1()(int,int){
//	return 1,2
//}
//func main() {
//	x,_:=test1();
//	println(x)
//}

func test3() (int, int) {
	return 1, 2
	//多个返回值可以直接作为其他函数调用实参
}
func add(x, y int) int {
	return x + y
}
func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}
func main() {
	println(add(test3()))
	println(sum(test3()))
}
