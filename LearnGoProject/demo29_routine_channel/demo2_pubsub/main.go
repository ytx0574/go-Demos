package main

import (
	"context"
	"fmt"
	"sync"
)

//todo: https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch1-basic/ch1-06-goroutine.html

//todo godoc的vfs包下面的gatefs, 用于控制虚拟文件系统的最大并发数
//https://github.com/golang/tools/blob/v0.1.0/godoc/vfs/gatefs/gatefs.go
//内部使用gatefs来包装原有的虚拟文件系统, 每一次得操作都是对chan得读写, 进而达到控制并发量得问题

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
				fmt.Printf("gggg : %v\n", i)
			}
		}
	}()
	return ch
}
// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <-chan int, prime, idx int) chan int {
	out := make(chan int)
	ff := func() {
		for {
			i := <-in
			if i % prime != 0 {
				select {
				case <- ctx.Done():
					return
				case out <- i:
					fmt.Printf("block <- :%v\n", i)
				}
			}
		}
	}
	go ff()

	return out
}

//todo 求素数 使用ctx来保证线程安全
//func demo1-mysql() {
//ctx, cancel := context.WithCancel(context.Background())
//ch := GenerateNatural(ctx) // 自然数序列: 2, 3, 4, ...
//for i := 0; i < 10; i++ {
//prime := <-ch // 新出现的素数
//fmt.Printf(">>>>%v: %v\n", i, prime)
////todo Prime内部一直读取, 本质上不会阻塞.  但是Prime内部还有一个无缓冲的chan, 导致堵塞
//ch = PrimeFilter(ctx, ch, prime, i) // 基于新素数构造的过滤器
//}
//cancel()
//}


func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hellll")
		case <- ctx.Done():
			//todo 通过ctx来实现安全退出或超时  此处会输出具体原因: 取消/超时
			fmt.Printf("ctx err:%v\n", ctx.Err())
			return ctx.Err()
		}
	}
}

func main() {

	////todo 使用ctx cancel来做超时或收尾工作
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go worker(ctx, &wg)
	//}
	//
	//time.Sleep(time.Second * 1)
	//cancel()
	//wg.Wait()


	fmt.Println("\xe4\xb8\x96") // 打印: 世
	fmt.Println("\xe7\x95\x8c") // 打印: 界
	fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc

	for _, i := range "我觉11" {
		fmt.Println(i)
	}

	const s = "\xe4\x00\x00\xe7\x95\x8cabc"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %x\n", i, s[i])
	}



	//p := models.NewPublisher(time.Millisecond*100, 5)
	//defer p.Close()
	//
	//all := p.Subscribe()
	//golang := p.SubscribeTopic(func(v interface{}) bool {
	//	if s, ok := v.(string); ok {
	//		return strings.Contains(s, "golang")
	//	}
	//	return false
	//})
	//
	//p.Publish("hello, world")
	//p.Publish("hello, golang")
	//
	//go func() {
	//	for msg := range all {
	//		fmt.Println("all >>", msg)
	//	}
	//}()
	//
	//go func() {
	//	for msg := range golang {
	//		fmt.Println("golang >>", msg)
	//	}
	//}()
	//
	//time.Sleep(time.Second)
	//fmt.Println("--")
}