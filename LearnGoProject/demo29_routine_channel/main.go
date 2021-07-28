package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

/*
go goroutine和channel
1. channel为引用类型, 使用时需要像map一样, 先初始化.  channel的make和slice/map不一样. channel 只能有两个参数, 第二个参数为cap. 不是size
*/

var (
	sliceLen int = 1000000000  //此处的数值越大 waitGroup和channel反而耗时更长了
	separatorItemCount = 10000
	mutex = sync.Mutex{}
)

func test() {
	slice := make([]int, sliceLen)
	maxCount := sliceLen / separatorItemCount
	fmt.Printf("separatorItemCount:%d, maxCount:%d\n", separatorItemCount, maxCount)

	timeBegin := time.Now()
	for count := 0; count < maxCount; count++ {
		//fmt.Printf("count:%d\n", count)
		var func1 = func (count int) {
			for i := count * separatorItemCount; i < count * separatorItemCount + separatorItemCount; i++ {
				//fmt.Printf("count:%d, i:%d, sliceMaxValue:%d\n", count, i, sliceMaxValue)
				slice[i] = i
			}
		}
		func1(count)
	}
	timeEnd := time.Now()
	timeInterval := timeEnd.Sub(timeBegin)

	//writeToFile("/Users/johnson/Desktop/slice.txt", slice)

	fmt.Printf("begin:%v, end:%v, 时间间隔:%v\n", timeBegin, timeEnd, timeInterval)
	fmt.Printf("slice长度为:%d, slice最后十个数为:%v\n", len(slice), slice[len(slice) - 10 :])
}

//使用同步锁. 这个应该是最慢的
func test_mutex() {
	slice := make([]int, sliceLen)
	intChan := make(chan int, 1000)
	fmt.Printf("intChanLen:%d\n", cap(intChan))
	maxCount := sliceLen / separatorItemCount
	fmt.Printf("separatorItemCount:%d, maxCount:%d\n", separatorItemCount, maxCount)

	timeBegin := time.Now()
	for count := 0; count < maxCount; count++ {
		//fmt.Printf("count:%d\n", count)
		var func1 = func (count int, intChan chan int) {
			for i := count * separatorItemCount; i < count * separatorItemCount + separatorItemCount; i++ {
				//fmt.Printf("count:%d, i:%d, sliceMaxValue:%d\n", count, i, sliceMaxValue)
				mutex.Lock()
				slice[i] = i
				mutex.Unlock()
			}
		}
		go func1(count, intChan)
	}

	time.Sleep(time.Second * 5)
	timeEnd := time.Now()
	timeInterval := timeEnd.Sub(timeBegin)


	fmt.Printf("begin:%v, end:%v, 时间间隔:%v\n", timeBegin, timeEnd, timeInterval)
	fmt.Printf("slice长度为:%d, slice最后十个数为:%v\n", len(slice), slice[len(slice) - 10 :])
}

//使用同步锁. 这个应该是最慢的
func test_waitGroup() {
	slice := make([]int, sliceLen)

	maxCount := sliceLen / separatorItemCount
	fmt.Printf("separatorItemCount:%d, maxCount:%d\n", separatorItemCount, maxCount)
	waitGroup := sync.WaitGroup{}

	timeBegin := time.Now()
	for count := 0; count < maxCount; count++ {
		//fmt.Printf("count:%d\n", count)
		var func1 = func (count int) {
			for i := count * separatorItemCount; i < count * separatorItemCount + separatorItemCount; i++ {
				//fmt.Printf("count:%d, i:%d, sliceMaxValue:%d\n", count, i, sliceMaxValue)
				slice[i] = i
			}
			waitGroup.Done()
		}
		waitGroup.Add(1)
		go func1(count)
	}

	waitGroup.Wait()

	timeEnd := time.Now()
	timeInterval := timeEnd.Sub(timeBegin)

	fmt.Printf("begin:%v, end:%v, 时间间隔:%v\n", timeBegin, timeEnd, timeInterval)
	fmt.Printf("slice长度为:%d, slice最后十个数为:%v\n", len(slice), slice[len(slice) - 10 :])
}

//此处使用goroutine
func test_channel() {
	//intChan := make(chan int, sliceLen)
	slice := make([]int, sliceLen)

	maxCount := sliceLen / separatorItemCount
	fmt.Printf("separatorItemCount:%d, maxCount:%d\n", separatorItemCount, maxCount)
	exitChan := make(chan bool, maxCount)
	fmt.Printf("exitChan cap:%d\n", cap(exitChan))


	timeBegin := time.Now()
	for count := 0; count < maxCount; count++ {
		var func1 = func (count int) {
			for i := count * separatorItemCount; i < count * separatorItemCount + separatorItemCount; i++ {
				//intChan<- i    //channel是线程安全的, 不如直接存效率高
				slice[i] = i
			}
			exitChan <- true
		}
		go func1(count)
	}

	//labelGetIntChan:
	//for {
	//	select {
	//	case v := <-intChan:
	//		slice[v] = v
	//		//slice = append(slice, v)
	//	default:
	//		fmt.Println("管道中没有值了\n")
	//		break labelGetIntChan
	//	}
	//}

	for count := 0; count < maxCount; count++ {
		<-exitChan
	}
	//close(exitChan) //这里关不关都可以. 因为exitChan中的值已经取完
	//close(intChan)
	//go func() {
	//	for i := range intChan {
	//		slice[i]  = i
	//	}
	//}()

	timeEnd := time.Now()
	timeInterval := timeEnd.Sub(timeBegin)

	//writeToFile("/Users/johnson/Desktop/slice2.txt", slice)

	fmt.Printf("begin:%v, end:%v, 时间间隔:%v\n", timeBegin, timeEnd, timeInterval)
	fmt.Printf("slice长度为:%d, slice最后十个数为:%v\n", len(slice), slice[len(slice) - 10 :])
}


func writeToFile(filePath string, slice []int) {

	file, _ := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, 0777)
	defer file.Close()
	bytes, _ := json.Marshal(slice)
	file.Write(bytes)
}


func main ()  {


	cpuNum := runtime.NumCPU()
	fmt.Printf("cpu num: %d\n", cpuNum)
	runtime.GOMAXPROCS(cpuNum)

	test()
	//test_mutex()
	test_waitGroup()
	test_channel()

	//var str = ""
	//for i := 0; i < 1000; i++ {
	//	sub := strconv.Itoa(i)
	//	if len(sub) == 1 {
	//		str += "00" + sub + ","
	//	}else  if len(sub) == 2 {
	//		str += "0" + sub + ","
	//	}else {
	//		str += "" + sub + ","
	//	}
	//}
	//fmt.Println(str)
}



