package double_list

import (
	"fmt"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list"
)

func AddNode(head *linked_list.Node, list *linked_list.Node) {
	temp := head
	isExsit := false
	for {
		if temp.Next == nil {
			break
		}else if temp.Next == list {
			isExsit = true
			fmt.Printf("该数据已存在链表中...\n")
			break
		}
		temp = temp.Next
	}

	if !isExsit {
		list.Pre = temp
		temp.Next = list
	}
}

func AddSortNode(head *linked_list.Node, list *linked_list.Node) {
	temp := head
	isExsit := false
	for  {
		if temp.Next == nil {
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
		list.Pre = temp

		//原来本位置的节点的前节点替换
		if temp.Next != nil {
			temp.Next.Pre = list
		}
		temp.Next = list
	}
}

func DelNode(head *linked_list.Node, id int) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}else if temp.Next.Id == id {
			//删除下一个节点
			//替换下一个节点
			temp.Next = temp.Next.Next
			//替换下一个节点的前节点
			temp.Next.Pre = temp
		}else {
			temp = temp.Next
		}
	}
}

func UpdateNode(head *linked_list.Node, list *linked_list.Node) {

	temp := head
	for {
		if temp.Next == nil {
			break
		}else if temp.Next.Id == list.Id && temp.Next.Name == list.Name {
			//更新list的前节点
			list.Pre = temp.Next.Pre
			//更新list的后节点
			list.Next = temp.Next.Next
			//更新后一个节点的上一个节点
			if temp.Next.Next != nil {
				temp.Next.Next.Pre = list
			}
			temp.Next = list
			break
		}else {
			temp = temp.Next
		}
	}
}

func ShowNode(head *linked_list.Node) {
	temp := head
	fmt.Printf("从左到右>>>   ")
	for {
		fmt.Printf("[%v %v %v]>>", temp.Id, temp.Name, temp.NickName)

		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	fmt.Println()

	fmt.Printf("从右到左<<<   ")
	for {
		fmt.Printf("[%v %v %v]>>", temp.Id, temp.Name, temp.NickName)

		if temp.Pre == nil {
			break
		}
		temp = temp.Pre
	}
	fmt.Println()
}

func CleanNode(head *linked_list.Node)  {
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		t := temp.Next
		temp.Next = nil
		temp.Pre = nil
		temp = t
	}
}