package main

//和z必须和方法在用一个作用于，才能被隐士返回，否则必须显示
func add2(x, y int) (z int) {
	z = x + y
	return

}
func main() {
	println(add2(1, 2))
}

//defer 延迟调用通过闭包都读取和修改
