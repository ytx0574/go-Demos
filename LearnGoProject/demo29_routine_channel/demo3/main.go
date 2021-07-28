package main

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan<- i
		fmt.Printf("写入数据:%v\n", i)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	fmt.Printf("read before\n")
	for {
		v, ok := <-intChan
		if !ok {
			break
		}

		exitChan<- false
		fmt.Printf("读取数据:%v\n", v)
	}
	fmt.Printf("read end\n")
	exitChan<- true
	close(exitChan)
}

func main() {
	var intChan chan int = make(chan int, 100)
	var exitChan chan bool = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//上面两行注释任意一行, 都会导致exitChan无法close. 而在主线程循环获取没有close的chan, 会导致死锁
	//仅开启第二行 出现死锁是因为, 协程里面获取一个未close的管道, 导致代码无法继续执行, 所以下面的exitChan也就无法close
	var i int
	for  {

		v, ok := <- exitChan
		if !ok {
			break
		}
		i++
		fmt.Printf("获取到exit:%v, i:%d\n", v, i)
	}

}