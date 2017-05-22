package main

import (
	"fmt"
	//"math/rand"
	"sync"
	//"time"
)

//
//func main() {
//	t:=NewTest()
//	println(<-t)//等待结束返回
//
//}
//func NewTest()chan int {
//	//简单工厂模式打包并发任务和channel
//	c:=make(chan int)
//
//	rand.Seed(time.Now().UnixNano())
//
//	go func() {
//		time.Sleep(time.Second)
//		c<-rand.Int()
//	}()
//	return c
//
//}
func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)

	sem := make(chan int, 1)

	for i := 0; i < 3; i++ {
		//println(i,"111111111111111")
		go func(id int) {
			//println(i,"22222222222")
			defer wg.Done()
			sem <- 1 //向sem发送数据，阻塞或者成功

			for x := 0; x < 3; x++ {
				fmt.Println(id, x)
				//println(x,"xxxxxxxxxxxxxxxxx")
			}
			<-sem //接收数据，使得其他阻塞goroutine可以发送数据
		}(i)
	}
	wg.Wait()
}
