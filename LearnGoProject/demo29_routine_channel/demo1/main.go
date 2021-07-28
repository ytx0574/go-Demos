package main

import (
	"fmt"
	"sync"
)

// 需求:现在要计算 1-30 的各个数的阶乘，并且把各个数的阶乘放入到 map 中。
// 最后显示出来。要求使用 goroutine 完成

//使用锁来对myMap加锁 实现数据写入
var myMap = make(map[int]int)
var mutex = sync.Mutex{}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mutex.Lock()
	myMap[n] = res
	mutex.Unlock()
}

func main() {
	for i := 0; i < 30; i++ {
		go  test(i)
	}

	for i, v := range myMap {
		fmt.Printf("myMap[%d] = %v\n", i, v)
	}
}