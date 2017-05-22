package main

//defer延迟调用，在return的时候才会调用，如果多个defer则按栈的顺序调用
func test6(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		println(100 / x)
	}()
	defer println("c")

}

func main() {
	test7()
}

func test7() {
	x, y := 10, 20

	defer func(i int) {
		println("defer", i, y)
		//x被复制
	}(x)
	x += 10
	y += 100
	println("x=", x, "y=", y)
}
