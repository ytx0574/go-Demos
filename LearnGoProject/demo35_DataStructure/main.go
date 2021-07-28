package main

import (
	"fmt"
	"go-Demos/LearnGoProject/demo35_DataStructure/binary_search_tree"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list/cycle_list"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list/double_list"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list/josuphu"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list/single_list"
	queue2 "go-Demos/LearnGoProject/demo35_DataStructure/queue"
	"go-Demos/LearnGoProject/demo35_DataStructure/sparse_matrix"
)

func main() {
	sparse_matrix.SparseMatrix()

	queue := queue2.NewQueue(10)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.ShowQueue()
	queue.Pop()
	queue.Pop()
	queue.Pop()
	queue.Pop()

	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.ShowQueue()


	cycleQueue := queue2.NewCycleQueue(5)
	cycleQueue.Push(1)
	cycleQueue.Push(2)
	cycleQueue.Push(3)
	cycleQueue.Push(4)
	cycleQueue.Push(5)
	cycleQueue.Push(6)
	cycleQueue.ShowQueue()
	cycleQueue.Pop()
	cycleQueue.Pop()
	cycleQueue.Pop()
	cycleQueue.ShowQueue()
	cycleQueue.Push(6)
	cycleQueue.Push(7)
	cycleQueue.Push(8)
	cycleQueue.Push(9)
	cycleQueue.ShowQueue()


	head := linked_list.Node{}

	songjiang := linked_list.Node{
		Id: 1,
		Name: "宋江",
		NickName: "及时雨",
	}
	lujunyi := linked_list.Node{
		Id: 2,
		Name: "卢俊义",
		NickName: "玉麒麟",
	}
	lingchong := linked_list.Node{
		Id: 2,
		Name: "林冲",
		NickName: "豹子头",
	}
	wuyong := linked_list.Node{
		Id: 3,
		Name: "吴用",
		NickName: "智多星",
	}

	single_list.AddNode(&head, &songjiang)
	single_list.AddNode(&head, &lujunyi)
	single_list.AddNode(&head, &lingchong)
	single_list.AddNode(&head, &wuyong)

	single_list.AddSortNode(&head, &songjiang)
	single_list.AddSortNode(&head, &wuyong)
	single_list.AddSortNode(&head, &lingchong)
	single_list.AddSortNode(&head, &lujunyi)
	single_list.AddSortNode(&head, &lujunyi)
	single_list.ShowNode(&head)

	single_list.DelNode(&head, 3)
	single_list.ShowNode(&head)

	lingchong2 := linked_list.Node{
		Id: 2,
		Name: "林冲",
		NickName: "豹子头2",
	}
	single_list.UpdateNode(&head, &lingchong2)
	single_list.ShowNode(&head)
	single_list.CleanNode(&head)
	single_list.ShowNode(&head)



	double_list.AddNode(&head, &songjiang)
	double_list.AddNode(&head, &lujunyi)
	double_list.AddNode(&head, &lingchong)
	double_list.AddNode(&head, &wuyong)
	double_list.AddNode(&head, &lingchong)
	double_list.AddNode(&head, &wuyong)
	double_list.ShowNode(&head)

	double_list.DelNode(&head, 1)
	double_list.ShowNode(&head)

	double_list.UpdateNode(&head, &lingchong2)
	double_list.ShowNode(&head)
	double_list.CleanNode(&head)
	double_list.ShowNode(&head)

	cycle_list.ShowNode(&head)
	cycle_list.AddNode(&head, &songjiang)
	cycle_list.AddNode(&head, &lujunyi)
	cycle_list.AddNode(&head, &lujunyi)
	cycle_list.AddNode(&head, &lingchong)
	cycle_list.AddNode(&head, &wuyong)
	cycle_list.ShowNode(&head)

	//cycle_list.DelNode(&head, 3)
	//cycle_list.ShowNode(&head)

	cycle_list.UpdateNode(&head,&lingchong2)
	cycle_list.ShowNode(&head)

	fmt.Printf("head是否闭环:%v\n", cycle_list.HasCycle(&songjiang))


	//linked_list.TestBankServiceCycleQueue()
	//linked_list.TestBankServiceCycleList()

	first := josuphu.AddBoy(5)
	josuphu.ShowBoys(first)
	josuphu.PlayGame(first, 3, 2)

	rootNode := binary_search_tree.BinaryTree()
	binary_search_tree.InfixOrder(rootNode)
	binary_search_tree.PreOrder(rootNode)
	binary_search_tree.SufOrder(rootNode)


	//str := ""
	//for i := 0; i< 1000; i++ {
	//	if i < 10 {
	//		str += "00" + strconv.Itoa(i)
	//	}else if i < 100 {
	//		str += "0" + strconv.Itoa(i)
	//	}else {
	//		str += "" + strconv.Itoa(i)
	//	}
	//	str += ","
	//}
	//fmt.Println(str)
}