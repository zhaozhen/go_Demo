package main

type Tester interface {
	Do()
}

type FuncDo func()

func (self FuncDo) Do() {
	self()
}
func main() {
	var t Tester = FuncDo(func() { println("hello,World!") })
	t.Do()
}
