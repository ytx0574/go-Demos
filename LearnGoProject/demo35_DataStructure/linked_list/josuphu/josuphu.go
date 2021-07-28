package josuphu

import "fmt"

/*
约瑟夫环:  小孩子丢手绢问题, 从谁开始, 每次数几个数, 数到的那个人出列

重点在于: 优先获得环第一个和最后一个, 这么一来, 尾部的那个next永远都是头部那个, 这样就可以从环中丢弃数到的那个人(tail.next = head.next), head丢弃
*/

type Boy struct {
	No int
	Next *Boy
}


func AddBoy(count int) *Boy {

	if count == 0 {
		 panic("最少要有一个孩子")
	}

	var first *Boy
	var current *Boy

	for i := 1; i <= count; i++ {

		boy := &Boy{No: i}

		if i == 1 {
			first =	boy
			current = boy
			current.Next = boy
		}else {
			current.Next = boy
			current = boy
			current.Next = first
		}
	}

	return first
}

func ShowBoys(first *Boy) {

	temp := first

	for {
		fmt.Printf("boy %v\n", temp.No)
		if temp.Next == first {
			break
		}
		temp = temp.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int) {

	current := first
	//先获取最后一个孩子
	tail := first
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	//先走到开始数数的数值
	for i := 1; i <= startNo - 1; i++ {
		current = current.Next
		tail = tail.Next
	}

	//开始数数, 出列
	for {
		//出列时注意, 出列的是最后数据的那个人
		for i := 1; i <= countNum - 1; i++ {
			current = current.Next
			tail = tail.Next
		}
		fmt.Printf("%v出队列了\n", current.No)

		//只有一个元素时 出队列
		if current.Next == current {
			break
		}

		//把当前元素出列
		tail.Next = current.Next
		//当前指向下一个元素
		current = current.Next
	}
}
