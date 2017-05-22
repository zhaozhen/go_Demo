package main

import (
	//"math"
	"runtime"
	"sync"
)

//func main() {
//	wg:=new(sync.WaitGroup)
//	wg.Add(4)
//
//	for i:=0;i<4;i++ {
//		go func(id int){
//			defer wg.Done()
//			sum(id)
//		}(i)
//	}
//	wg.Wait()
//
//}
//func sum(id int)  {
//	var x int64
//	for i:=0;i< math.MaxUint32;i++ {
//		x+=int64(i)
//	}
//
//	println(id,x)
//}
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 6; i++ {
			println(i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		println("Hello,World!")
	}()

	wg.Wait()
}
