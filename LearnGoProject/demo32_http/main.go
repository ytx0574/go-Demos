package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"unsafe"
)
var ctx context.Context = context.Background()

/*
context的实践
*/

func init () {
	fmt.Printf("init 1\n")
}

func init()  {
	fmt.Printf("init 2\n")
}

func main () {
	TestUintptr()
	//todo 前三者都返回一个子context 和 CancelFunc
	newCtx, newFunc := context.WithTimeout(ctx, time.Second)
	defer newFunc()
	go HelloHandle(newCtx, time.Millisecond * 500)
	select {
	case <- newCtx.Done():
		fmt.Printf("Hello Handle %v\n", ctx.Err())
	}

	return

	newCtx, newFunc = context.WithCancel(ctx)
	fmt.Printf("ctx:%s, newCtx:%s\n", ctx, newCtx)
	defer newFunc()
	go Speak(newCtx)  //带入的参数是有效的子context, 不应该用ctx
	time.Sleep(time.Second * 5)
	fmt.Printf("ctx Canceled!!!\n")
	newFunc() //主动取消

	newCtx, newFunc = context.WithDeadline(ctx, time.Now().Add(time.Second))
	defer newFunc()
	go Monitor(newCtx)
	time.Sleep(time.Second * 3)


	//todo 生成一个子context, 绑定带入的key value  所有派生的context都可以获取值, 如key相同, 优先获取自己
	subCtx := context.WithValue(ctx, "key", "value")
	fmt.Printf("%v\n", subCtx.Value("key"))


	subCtx2 := context.WithValue(subCtx, "key", "value2")
	fmt.Printf("%v\n", subCtx2.Value("key"))
	fmt.Printf("%v\n", subCtx.Value("key"))
}

func HelloHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Printf("ctx done:%v\n", ctx.Err())
	case <-time.After(duration):
		fmt.Printf("after %v\n", duration)
	}
}

func Speak(ctx context.Context) {
	timeTick:
	for range time.Tick(time.Second * 2) {
		select {
		case <-ctx.Done():
			fmt.Printf("说话退出 %v\n", ctx.Err())
			break timeTick
		default:
			fmt.Printf("一直输出说话 %v\n", ctx.Err())
		}
	}
}

func Monitor(ctx context.Context)  {
	select {
	case <-ctx.Done(): //到了时间, 执行结束
		fmt.Printf("ctx done:%v\n", ctx.Err())
	case <-time.After(time.Second):
		fmt.Printf("stop monitor %v\n", ctx.Err())
	}
}


func TestContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func TestContextCancel2() {
	fmt.Printf("start work...\n")

	newContext, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		for  {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				//如主线程没有其他代码, 这里也不会执行. 因为主线程已经执行完了, 关闭了进程.
				fmt.Printf("work done... %v\n", ctx.Err())
				return  //此处不用使用break. 无法跳出for  或者使用label跳出也行
			default:
				fmt.Printf("working... %v\n", ctx.Err())
			}
		}
	}(newContext)
	time.Sleep(time.Second * 3)
	cancel()

	fmt.Printf("end work...\n")
	// time.Sleep(time.Second * 5)
	//fmt.Printf("end work2...\n")
}

type data struct {
	x [1024 * 100]byte
}

var Condition bool = false

func TestUintptr() {
	const N = 100
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		var ptr uintptr = uintptr(unsafe.Pointer(&data{}))
		cache[i] = ptr
	}

	m := sync.Mutex{}
	con := sync.NewCond(&m)
	go func() {
		//con.Wait()
		//con.Signal()
		//con.Broadcast()
		con.L.Lock()
		for !Condition {
			con.Wait()
		}
		fmt.Printf("--收到信号, 继续执行\n")
		con.L.Unlock()
	}()

	//con.L.Lock()  //这里加不加锁都可以
	time.Sleep(time.Second * 2)
	fmt.Println("demo1-mysql goroutine ready")
	Condition = true
	con.Signal()
	fmt.Println("demo1-mysql goroutine broadcast")
	//con.L.Unlock()
	fmt.Printf("-----\n")

}
