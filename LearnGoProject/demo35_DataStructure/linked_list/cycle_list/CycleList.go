package cycle_list

import (
	"fmt"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list"
)

func AddNode(head *linked_list.Node, list *linked_list.Node) {
	if head.Next == nil {
		head.Id = list.Id
		head.Name = list.Name
		head.NickName = list.NickName
		head.Next = head
		return
	}

	temp := head
	isExist := false
	for {
		if temp.Next == list {
			fmt.Printf("不允许添加重复\n")
			isExist = true
			break
		}else if temp.Next == head {
			break
		}
		temp = temp.Next
	}

	if !isExist {
		list.Next = head
		temp.Next = list
	}
}

func AddSortNode(head *linked_list.Node, list *linked_list.Node) {
	if head.Next == nil {
		AddNode(head, list)
	}


	temp := head
	isExsit := false
	for  {
		if temp.Next == head {
			break
		}else if (temp.Next == list) {
			fmt.Printf("该数据已存在链表中...\n")
			isExsit = true
			break
		}else if temp.Next.Id > list.Id {
			break
		}
		temp = temp.Next
	}

	if !isExsit {
		//插入节点的前后
		list.Next = temp.Next
		temp.Next = list
	}
}

func DelNode(head *linked_list.Node, id int) {
	temp := head
	for temp.Next != nil {
		if temp.Next == head {
			break
		}else if temp.Next.Id == id {
			temp.Next = temp.Next.Next
		}else {
			temp = temp.Next
		}
	}
}

func UpdateNode(head *linked_list.Node, list *linked_list.Node) {
	temp := head
	for temp.Next != nil {
		if temp.Next == head {
			break
		}else if temp.Next.Id == list.Id && temp.Next.Name == list.Name {
			list.Next = temp.Next.Next
			temp.Next = list
		}
		temp = temp.Next
	}
}

func HasCycle(head *linked_list.Node) bool  {
	if head == nil || head.Next == nil {
		return false
	}
	//todo:闭环. faster和slower, faster每次跨两步, slower每次跨一步.
	//todo:比如跑步, 快的那个始终都会追上慢的那个
	faster := head
	slower := head
	for faster != nil && slower != nil {
		faster = faster.Next.Next
		slower = slower.Next
		if faster == slower {
			return true
		}
	}
	return false
}

func ShowNode(head *linked_list.Node) {
	fmt.Printf("环形链>>>   ")
	temp := head
	for temp.Next != nil {
		fmt.Printf("[%v %v %v]>>", temp.Id, temp.Name, temp.NickName)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	fmt.Println()
}
