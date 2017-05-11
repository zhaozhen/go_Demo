package main

import "fmt"

func main() {
	//test8()
	test9()
}

//没有结构化异常，使用panic抛出错误，recover接受错误

func test8()  {

	defer func(){
		if err:=recover();err!=nil{
			println(err.(string))
			println("0-----------")
		}
	}()

	panic("panic error")
}
func test9()  {
	defer func(){
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic-----zuihou")
	}()

	panic("test panic")
}
//