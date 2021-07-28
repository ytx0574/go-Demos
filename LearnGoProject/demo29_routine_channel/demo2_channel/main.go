package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo29_routine_channel/demo2_channel/utils"
	_ "go-Demos/LearnGoProject/demo29_routine_channel/demo2_channel/utils2"
	"math/rand"
	"reflect"
	"sync"
	"time"
)
/*
channel 说明:
1. 在没有使用协程得情况下, 超读或超写管道, 会报告deadlock错误. 如果在协程里面, 超读或超写, 不会报错, 但是代码后面不会继续执行. 管道数据取出后可以再次放入
2. 管道是引用类型, 使用前必须使用make初始化. make的第二个参数为cap. 不是size
3. 管道的数据是先入先出 First In First Out, 本质是一个队列
4. 线程安全, 多goruntine访问时, 不需要加锁.
5. channel是有类型的, 必须放指定类型
6. 管道内部参数为空接口类型时, 需注意使用类型断言转换回来 虽然运行时知道它是什么, 但是编译器编译时不知道
7. todo 遍历管道同第一条一致, 非协程状态遍历时(for range), 需要先close, 否则死锁. 如只遍历已知条数, 不需要close即可遍历;
    协程里面遍历可以不用close. 但是数据获取完之后, 后面也不会继续执行. 遍历完成后, 管道内数据已清空. (如想继续执行后面内容, 则需提前使用close);
    主协程里面的死锁, 会导致崩溃, 其他协程里面的死锁没有任何提示, 仅仅是死锁后, 后续代码无法执行
8. 管道关闭并且已经取完值之后, 再次取值, 如果是引用类型, 返回nil. 非引用类型, 返回初始值. 管道没有关闭的时候, 再次取值会造成死锁. 常规情况就是写入多少, 取出多少
9. 管道可以声明为只读或只写. 最佳实践demo5
10. select可以在没有关闭管道时, 获取内部的内容(内部case只能操作channel"读 写 default"). 可以使用return或break终止.  如果是return, 后续也不再继续执行.
	select内部会尽量均匀的分配给每个case相同的执行时间. 通过select的case time.After可以很简单的实现超时效果.
11. goroutine中可以捕获异常,  但是捕获异常的前提是goroutine本身没有死锁, 本身不退出, 不会执行到recover
12. todo channel默认的创建, len/cap=0  即不带缓冲区, 只能先存后取, 再存再取(否则死锁报错), 带缓冲区的只会阻塞, 不会锁住.
13. 阻塞条件: 带缓冲区的, 写入阻塞条件: 缓冲区满, 读取阻塞条件:缓冲区为空;  不带缓冲区的, 读写阻塞条件: 同一时间, 没有另一线程与之对应的反向操作(你读我写 | 你写我读);
14. 可以使用close chan来广播使多个goroutine同时关闭, 配合waitGroup来保证关闭的同时, 能处理完所有的清理工作
*/

func main() {
	var intChan chan int
	fmt.Printf("intChan:%v, intChan地址:%p\n", intChan, &intChan)
	fmt.Printf("len intChan:%d, cap intChan:%d\n", len(intChan), cap(intChan))

	intChan = make(chan int, 3)

	fmt.Printf("intChan:%v, intChan地址:%p\n", intChan, &intChan)
	fmt.Printf("len intChan:%d, cap intChan:%d\n", len(intChan), cap(intChan))

	intChan<- 1
	intChan<- 2
	intChan<- 3
	fmt.Printf("len intChan:%d, cap intChan:%d\n", len(intChan), cap(intChan))

	go func() {
		fmt.Printf("执行1\n") //默认执行
		<-intChan

		intChan<- 4
		fmt.Printf("执行2\n") //执行, 因为intChan先移除一个, 然后再次补入一个.


		<-intChan
		<-intChan
		<-intChan
		fmt.Printf("取值:%v\n", <-intChan) //不执行. 因为三个已经被移除了, 后面的就不会执行

	}()

	//go func() {
	//	fmt.Printf("执行11\n") //默认执行
	//	intChan <- 4
	//	fmt.Printf("执行22\n") //默认执行
	//}()

	time.Sleep(time.Second)

	var allChan chan interface{} = make(chan interface{}, 5)
	allChan<- 6
	allChan<-utils.Cat{
		Name: "Tom",
		Age: 11,
	}
	allChan<-"333"

	<-allChan
	cat := <-allChan;
	//fmt.Printf("cat:%v\n", cat.Name)
	fmt.Printf("cat:%v\n", cat.(utils.Cat).Name)


	var catChan = make(chan utils.Cat, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		catChan<- utils.Cat{Name: "Tom~" + fmt.Sprintf("%d", i), Age: rand.Int() % 10 + 1}
	}

	//close(catChan)  //遍历channel时, 必须先close掉.
	//for v := range catChan {
	//	fmt.Printf("catName:%v, catAge:%d\n", v.Name, v.Age)
	//}

	go func() {
		//for {
		//	cat, ok := <-catChan
		//	if !ok {
		//		break
		//	}
		//	fmt.Printf("获取cat:%v\n", cat)
		//}
		for v := range catChan {
			fmt.Printf("catName:%v, catAge:%d\n", v.Name, v.Age)
		}
		fmt.Printf("只有close掉, 这里才会执行..., len catChan:%d\n", len(catChan))
		v := <-catChan
		v = <-catChan
		v = <-catChan
		//v.Name = "11"
		if reflect.DeepEqual(v, utils.Cat{}) {
			fmt.Printf("chan数据已读取完的情况下, 再次读取返回基本数据类型的初始值. 引用类型返回nil\n")
		}

		fmt.Printf("只有close掉, 这里才会执行..., len catChan:%d %q %q\n", len(catChan), v, utils.Cat{})
	}()


	time.Sleep(time.Second)
	selectTimeout()
}

//todo 使用select来实现timeout
func selectTimeout(){
	w := make(chan bool)
	c := make(chan int, 2)
	go func() {
		select {
		case v:=<-c:
			fmt.Println(v, "输出chan")
		case <-time.After(time.Second):
			fmt.Println("超时")
		}
		w <- true
	}()
	c <- 1 //todo 此段直接引发超时
	<- w
}

//todo 带缓冲的channel
func bufferChannel() {
	exit := make(chan  bool)
	//chs := make(chan  bool, 1)
	exit1 := make(chan  bool)
	fmt.Println("chs" ,len(exit), cap(exit))
	fmt.Println("chs1" ,len(exit1), cap(exit1))

	//带缓冲区
	ch := make(chan interface{}, 3)
	fmt.Println(len(ch))
	ch <- 1
	ch <- 2
	ch <- 3

	go func() {

		for a := range ch {
			fmt.Println(len(ch))
			fmt.Println(1111, a)
		}
		fmt.Println("45465")
		exit <- true
		exit <- true

		fmt.Println("454651")
	}()
	ch <- 4
	ch <- 5
	close(ch)  //todo 用完必须关闭ch, 达到关闭队列的作用. 否则上面的for range则会被阻塞, 导致后面的退出指令无法发送
	<-exit   //todo 等待关闭指令
	//<-chs
	//close(chs)
	if d, ok := <-exit; ok {
		fmt.Println(ok, ok)
	}else {
		fmt.Println(d, ok)
	}
}

//todo 使用channel实现信号量, 保证内部连续执行
func semaphore() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	sem := make(chan int, 1)
	for i := 0; i < 3; i ++ {
		go func(x int) {
			defer wg.Done()

			sem <- 1  //todo 向sem发送数据, 阻塞或成功
			for i:=0; i< 3; i++ {
				fmt.Println("哈哈哈", x)
			}
			<-sem  //todo 用完接受数据, 使得其它被阻塞的协程可以发送数据
		}(i)
	}
	wg.Wait()
}

//todo 使用close channnel来通知退出
func closeChannelQuit(){
	wg := sync.WaitGroup{}
	quit := make(chan bool)

	for i := 0;i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			task := func() {
				fmt.Println(i, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}

			for  {
				select {
				case a :=<- quit:
					fmt.Println("退出", i, a)   //todo close channel不会阻塞, 因此可用作退出通知
					return
				default:
					fmt.Println("执行任务", i)
					task()
				}
			}
		}(i)
	}

	time.Sleep(time.Second * 5)
	for i := 0;i < 5; i++ {
		quit<-true  //todo 此时仅某一个协程收到退出指令, 其他的协程还会继续执行, 或者自己知道多少个协程在监听, 自己发送多少次退出指令
	}
	//close(quit)  //todo 所有的协程都会收到退出指令
	wg.Wait()

	fmt.Println("继续执行...")
}