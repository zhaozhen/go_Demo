package main


func main()  {
	test4()
}
func test3() {
	var i int
	for{
		println(i)
		i++
		if i>2{ goto BREAK}
	}

BREAK:
	println("break")

}
func test4()  {
	//go的跳转感觉好高达上呀。。。。。
	L1:
	for x:=0;x<3;x++ {

		L2: for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}
			if x > 1 {
				continue L1
			}

			print(x, ":", y, " ")
		}
		println()
	}
}
//附记：break可以用与for，switch，select，而continue近能用于for循环