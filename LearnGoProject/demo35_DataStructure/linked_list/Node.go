package linked_list

type Node struct {
	Id int
	Name string
	NickName string
	Next *Node
	Pre *Node
}

type BinaryNode struct {
	 Val string
	 Left, Right *BinaryNode
}

//func (this *BinaryNode)String() {
//	fmt.Println(this.Val)
//}