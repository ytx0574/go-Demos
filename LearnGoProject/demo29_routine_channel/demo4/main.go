package main

import (
	"fmt"
	"time"
)

//要求统计 1-200000 的数字中，哪些是素数?这个问题在本章开篇就提出了，现在我们有 goroutine
//和 channel 的知识后，就可以完成了

func main()  {

	//获取素数开启的协程数
	var goroutineCount = 10
	var num = 200000

	begin := time.Now()

	intChan := make(chan int, 100)
	premeChan := make(chan int, 100)
	exitChan := make(chan bool, goroutineCount)

	go func() {
		for i := num; i >= 0; i-- {
			intChan<- i
		}
		close(intChan)
	}()

	for i := 0; i < goroutineCount; i++ {
		go func() {
			for  {
				v, ok := <-intChan
				if !ok {
					break
				}

				flag := true
				for i := 2; i < v; i++ {
					if (v % 2 == 0) {
						flag = false
						break
					}
				}
				if flag {
					premeChan<- v
				}
			}
			exitChan<- true
		}()
	}

	go func() {
		for  {
			//此处在已知exitChan里面有多少个变量时, 使用常规的for. 依次取出各位变量, 没有问题. 在正常取完的情况下会继续向下执行
			//但是使用for range或死讯获取一个没有关闭的管道时, 会导致后面代码无法执行, 致后面的premeChan出现死锁
			for i := 0; i < goroutineCount; i++ {
				<-exitChan
				fmt.Printf("exitChan\n")
			}
			//for v := range exitChan {
			//	fmt.Printf("exitChan %v\n", v)
			//}
			//for {
			//	_, ok := <-exitChan
			//	if !ok {
			//		break
			//	}
			//}
			close(premeChan)
		}
	}()

	slice := []int{}
	for {
		if v, ok := <-premeChan; ok {
			//fmt.Printf("素数:%v\n", v)
			slice = append(slice, v)
		}else {
			break
		}
	}
	end := time.Now()

	//fmt.Printf("slice:%v\n ", slice)
	fmt.Printf("获得耗时:%v\n", end.Sub(begin))


}



