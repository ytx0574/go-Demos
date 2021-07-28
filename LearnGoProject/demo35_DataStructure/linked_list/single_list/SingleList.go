package single_list

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
		list.Next = temp.Next
		temp.Next = list
	}
}

func DelNode(head *linked_list.Node, id int) {
	temp := head
	for {
		if temp.Next == nil {
			break
		}else if temp.Next.Id == id {
			//删除操作
			temp.Next = temp.Next.Next
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
			list.Next = temp.Next.Next
			temp.Next = list
			break
		}else {
			temp = temp.Next
		}
	}
}

func ShowNode(head *linked_list.Node) {
	temp := head
	for {
		fmt.Printf("[%v %v %v]>>", temp.Id, temp.Name, temp.NickName)

		if temp.Next == nil {
			break
		}
		temp = temp.Next
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
		temp = t
	}
}