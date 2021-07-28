package main

import (
	"fmt"
	"sync"
	"time"
)

func cal(a, b int, exitChan chan bool) {
	c := a + b
	fmt.Println("a + b = ", c)
	exitChan<- true
}
func cal1(a, b int, syncWait *sync.WaitGroup) {
	c := a + b
	fmt.Println("a + b = ", c)
	syncWait.Done()
}

func test(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("不带缓冲区的chan... set:%d\n", i)
		c <- i
	}
}

func main() {

	//常规的同步goroutine
	var exitChan chan bool = make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go cal(i, i + 1, exitChan)
	}

	for i := 0; i < 10; i++ {
		 <-exitChan
	}
	fmt.Println("执行完毕1")


	//syncWait 同步.    因为WaitGroup是结构体, 所以使用时, 需要使用引用传递i
	var syncWait = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		syncWait.Add(1)
		go cal1(i, i + 1, &syncWait)
	}
	syncWait.Wait()
	fmt.Println("执行完毕2")



	//默认创建的chan是不带缓冲区的, 只有一个值, 那么就是存取//存取 带cap的就是有缓冲区, 速度快的话, 可以存完 再取
	ch := make(chan int)
	go test(ch)
	for i := 0; i < 10; i++ {
		v := <- ch
		fmt.Printf("不带缓冲区的chan... get:%d\n", v)
	}


	//channel实现作业池
	//我们创建三个channel，一个channel用于接受任务，一个channel用于保持结果，还有个channel用于决定程序退出的时候。
	taskch := make(chan int, 20)
	resch := make(chan int, 20)
	exitChan1 := make(chan bool, 5)

	go func() {
		for i := 0; i < 10; i++ {
			taskch <- i
		}
		close(taskch)
	}()

	for i := 0; i < 10; i++ {
		go task(taskch, resch, exitChan1)
	}

	go func() {
		for i := 0; i < 10; i++ {
			<- exitChan1
		}
		close(resch)
		//t := <- exitChan1
		//fmt.Printf("超读exitChan1: %d\n", t)  超读, 下面的close无法执行.  如果把close(resch)也写在超读后面, 同样导致无法关闭   后面出现死锁
		close(exitChan1)  //此句不写也没问题, 因为上面的for没有超读
	}()

	for res := range resch {
		fmt.Printf("taks res:%v\n", res)
	}



	//在对channel进行读写的时，go还提供了非常人性化的操作，那就是对读写的频率控制，通过time.Ticker实现
	requests := make(chan int, 5)
	for i := 0; i < cap(requests); i++ {
		requests <- i
	}
	close(requests)

	//limiter := time.Tick(time.Millisecond * 500) //内部实现为NewTicker.C  返回一个只读的<-chan Time
	ticker := time.NewTicker(time.Millisecond * 500) //创建一个计时器. 可以阻塞协程.
	//for l := range limiter {
	//	fmt.Printf("llll:%v, time:%v\n", l, time.Now())
	//	break
	//}

	for i := range requests {
		 <-ticker.C
		fmt.Printf("控制获取频率, 输出i:%d\n", i)
		if i == 4 {
			ticker.Stop() //相当于关闭ticker中的 C (<-chan Time), 但是不可关闭后, 再次访问. 比如这里写成 i == 3
		}
	}

	fmt.Printf("time tick end\n")
}

func task(taskch, resch chan int, exitChan chan bool)  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("task()捕获到异常: %v\n", err)
		}
	}()

	for t := range taskch {
		fmt.Printf("do task %d\n", t)
		resch <- t  //提取任务结果
	}
	exitChan <- true
}