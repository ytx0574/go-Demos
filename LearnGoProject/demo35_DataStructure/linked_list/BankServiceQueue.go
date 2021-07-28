package linked_list

import (
	"fmt"
	"strconv"
	"time"
)

type BankService struct {
	id int
	customer string
	next *BankService
}

var ay []BankService
var head, rear int
var maxSize int

func InitQueue(num int) {
	maxSize = num
	ay = make([]BankService, num)

	for i, _ := range ay {
		ay[i].id = i
	}
}

func Push(name string) {
	if Full() {
		fmt.Printf("队列已满\n")
		return
	}
	fmt.Printf("%v 正在服务 %v\n", ay[rear].id, name)
	ay[rear].customer = name
	rear = (rear + 1) % maxSize
}

func Pop() {
	if Empty() {
		fmt.Printf("队列为空\n")
		return
	}

	fmt.Printf("%v 结束对 %v 的服务\n", ay[head].id, ay[head].customer)
	head = (head + 1) % maxSize
}

func Full() bool {
	return (rear+ 1) %maxSize == head
}

func Empty() bool {
	return rear == head
}



func TestBankServiceCycleQueue() {

	var maxSize = 5
	var popChan chan bool  = make(chan bool , maxSize)
	InitQueue(maxSize)

	go func() {
		time.Sleep(time.Second * 1)
		for {
			Pop()
			popChan <- true

			time.Sleep(time.Millisecond * 100)
		}
	}()

	for i := 300; i < 333; i++ {
		if !Full() {
			Push(strconv.Itoa(i))
		}else {
			select {
			case <-popChan:
				Push(strconv.Itoa(i))
			}
		}
	}
}


func TestBankServiceCycleList() {
	var maxSize = 5
	var first *BankService
	var current *BankService
	var popChan chan *BankService = make(chan *BankService, maxSize)
	var startCustomerId = 300
	var customerCount = 333

	for i := startCustomerId; i < customerCount; i++ {
		if i < startCustomerId + maxSize { //提前创建一个maxSize个数据的链表
			bank := &BankService{
				id: i - startCustomerId,
			}
			if i == startCustomerId {
				first = bank
				first.next = first
				current = first
				go func() {
					for {
						time.Sleep(time.Second)
						fmt.Printf("%v 结束对 %v 的服务\n", current.id, current.customer)
						popChan<- current
						current = current.next
					}
				}()
			}else {
				current.next = bank
				bank.next = first
				current = bank
			}

			current.customer = strconv.Itoa(i)
			fmt.Printf("%v 正在对 %v 进行服务\n", current.id, current.customer)
		}else {
			select {
			case x := <-popChan:
				x.customer = strconv.Itoa(i)
				fmt.Printf("%v 正在对 %v 进行服务\n", x.id, x.customer)
			}
		}
	}
}

