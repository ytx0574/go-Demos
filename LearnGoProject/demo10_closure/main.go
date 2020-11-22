package main

import (
	"fmt"
	"time"
)
//本章内容来自: (部分也没搞清楚)  https://www.jianshu.com/p/faf7ef7fbcf8
func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v) //输出ccc  主协程执行完之后, 才执行func
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("-------------")

	for _, v := range s {
		func () {
			fmt.Println(v)   // 输出abc  每一次for都执行了sleep, 让for中定义的子协程有时间执行
		}()
		time.Sleep(time.Second * 1)
	}
	fmt.Println("-------------")

	for i, v := range  s {
		func () {
			fmt.Println(v) //输出abc, 执行第二次时, sleep一会儿, 此时就定义好两个子协程执行
		}()
		if i == 1 {
			time.Sleep(time.Second * 1)
		}
	}
	fmt.Println("-------------")

	for _, v := range  s {
		func (v string) {
			fmt.Println(v)
		}(v)
	}
}