package main

import "fmt"

/**简单递归demo*/
func fbn(i int) int {
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fbn(i - 1) + fbn(i - 2)
	}
}

func f(n int) int {
	if n == 1 {
		return 3
	}else {
		return 2 * f(n - 1) + 1
	}
}

//递归获取桃子数,  每天吃一半并且多吃一个, 到了第十天的时候还有1个桃子.  获取前面几天的桃子数
func sum_peach(day int) int {
	if day == 10 {
		return 1
	}else {
		if day > 10 || day < 1 {
			return 0
		}
		return (sum_peach(day + 1) + 1) * 2
	}
}

func main () {
	fmt.Printf("fbn(1) = %v\n", fbn(1))
	fmt.Printf("fbn(3) = %v\n", fbn(2))
	fmt.Printf("fbn(7) = %v\n", fbn(7))


	fmt.Printf("f(1) = %v\n", f(1))
	//fmt.Printf("f(0) = %v\n", f(0))
	fmt.Printf("f(5) = %v\n", f(5))


	fmt.Printf("peach(10) = %v\n", sum_peach(11))
}