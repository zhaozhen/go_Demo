package main

import "fmt"

func main() {
	data := make(chan int, 3) //数据交换队列
	exit := make(chan bool)   //推出通知

	data <- 1 //发送数据
	data <- 2
	data <- 3

	go func() {
		for d := range data {
			//从队列接受数据，知道close
			fmt.Println(d)
		}

		fmt.Println("recv over")
		exit <- true //发出推出通知
	}()

	data <- 4
	data <- 5

	close(data) //关闭队列

	fmt.Println("send over")
	<-exit //等待推出通知
}
