package main

//匿名表达式
func main() {

	fn := func() {
		println("Hello,World!")
	}
	fn()
	fns :=[](func(x int) int){
		func(x int) int {return x+1},
		func(x int) int {return x+2},
	}
	println(fns[1](100))
//--------function as field-------------------------
	d:= struct {
		fn func() string
	}{
		fn:func() string {return "Hello,World!"},
	}
	println(d.fn())
	//-------channel fo function--------------------
	fc :=make(chan func() string,2)
	fc <- func() string{ return "hello,World!"}
	println((<-fc)())
}
