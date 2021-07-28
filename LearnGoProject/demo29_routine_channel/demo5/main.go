package main

import (
	"fmt"
)

func main() {
	//只读管道  不能关闭, 读取死锁
	//var onlyReadIntChan <-chan int = make(<-chan int, 10)
	//fmt.Println(onlyReadIntChan)
	//aa := <-onlyReadIntChan
	//fmt.Println(aa)

	//go func() {
	//	for {
	//		num := <-onlyReadIntChan
	//		fmt.Printf("%v  %v\n", num, aa)
	//	}
	//}()

	//只写管道, 可关闭. 写入超过cap死锁
	var onlyWriteIntChan chan<- int = make(chan<- int, 10)
	for i := 0; i < 10; i++ {
		onlyWriteIntChan<-i
	}
	fmt.Println("只读chan执行完")



	//只读 只写最佳实践
	var ch chan int = make(chan int, 10)
	var exitChan = make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束")

}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch<- i
	}
	close(ch)
	var a struct{}
	exitChan<- a
}

func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan<- a
}




