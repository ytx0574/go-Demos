package queue

import (
	"errors"
	"fmt"
)

/*
使用数组来实现队列 核心在于 (head + 1) % maxSize 取模运算, 来实现环形队列
*/

type Queue struct {
	MaxSize int
	head int
	tail int
	array []int
}

func NewQueue(maxSize int) Queue {
	return Queue{
		MaxSize: maxSize,
		head: 0,
		tail: 0,
		array: make([]int, maxSize),
	}
}

func (this *Queue)Push(val int) (err error) {
	if this.tail == this.MaxSize {
		return errors.New("队列已满")
	}
	this.array[this.tail] = val
	this.tail++
	return
}

func (this *Queue)Pop() (val int, err error) {
	if this.head == this.tail {
		return 0, errors.New("队列为空")
	}

	val = this.array[this.head]
	this.head++
	return
}

func (this *Queue)Size() int {
	return this.tail - this.head
}

func (this *Queue)ShowQueue() {
	fmt.Printf("队列元素个数为:%v\n", this.Size())
	for i := this.head; i < this.tail; i++ {
		fmt.Printf("%v\t", this.array[i])
	}
	fmt.Println()
}


type CycleQueue struct {
	Queue
}

func NewCycleQueue(maxSize int) (cycleQueue CycleQueue) {
	//利用数值的特殊性, 实际分配的大小为传入大小+1  留出一个空的容量来做计算
	//主要利用的取模的特性来实现循环
	cycleQueue = CycleQueue{
		Queue: Queue{
			MaxSize: maxSize + 1,
			array: make([]int, maxSize + 1),
		},
	}
	return
}

func (this *CycleQueue)Push(val int) (err error) {
	if this.IsFull() {
		err = errors.New("队列已满")
		fmt.Println(err)
		return
	}

	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.MaxSize
	return
}

func (this *CycleQueue)Pop() (val int, err error) {
	if this.IsEmpty() {
		err = errors.New("队列为空")
		fmt.Println(err)
		return 0, err
	}

	val = this.array[this.head]
	this.head = (this.head + 1) % this.MaxSize
	return
}

func (this *CycleQueue)IsFull() bool {
	return (this.tail + 1) % this.MaxSize == this.head
}

func (this *CycleQueue)IsEmpty() bool {
	return this.head == this.tail
}

func (this *CycleQueue)Size() int {
	//关键部分 在这里.  计算队列个数
	return (this.tail + this.MaxSize - this.head) % this.MaxSize
}
func (this *CycleQueue)ShowQueue() {
	fmt.Printf("队列元素个数为:%v\n", this.Size())

	head := this.head
	for i := 0; i < this.Size(); i++ {
		fmt.Printf("array[%v] = %v\t", i, this.array[head])
		head = (head + 1) % this.MaxSize
	}
	fmt.Println()
}

