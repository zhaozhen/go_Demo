package main

func main() {
	test2()
}
func test() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		println("a")
		//继续下一分支，并且不再判断条件
		fallthrough
	case 1, 3:
		println("b")
	default:
		println("c")

	}
}
func test2() {
	x := []int{1, 2, 3}
	switch i := x[2]; {
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")

	}
}
