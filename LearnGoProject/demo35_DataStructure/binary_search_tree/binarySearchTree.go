package binary_search_tree

import (
	"fmt"
	"go-Demos/LearnGoProject/demo35_DataStructure/linked_list"
)

//二叉查找树特性:
//所有左侧子节点都比根节点小, 所有右侧子节点都比根节点大
//https://zhuanlan.zhihu.com/p/37470148
func BinaryTree() (root *linked_list.BinaryNode) {
//			A
//	B			 	C
//D		E		F		G
//	H

	begin := int(byte('A'))
	end := int(byte('H'))
	sliceNode := make([]*linked_list.BinaryNode, end - begin + 1)
	for i := 0; i < len(sliceNode); i++ {

		sliceNode[i] = &linked_list.BinaryNode{
			Val: string([]byte{byte(i + begin)}),
		}
		fmt.Println(sliceNode[i])
	}

	A := sliceNode[0]
	B := sliceNode[1]
	C := sliceNode[2]
	D := sliceNode[3]
	E := sliceNode[4]
	F := sliceNode[5]
	G := sliceNode[6]
	H := sliceNode[7]

	A.Left = B
	A.Right = C

	B.Left = D
	B.Right = E

	D.Right = H

	C.Left = F
	C.Right = G

	return A
}

//①、中序遍历:左子树——》根节点——》右子树
func InfixOrder(node *linked_list.BinaryNode) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("中序遍历:%v\n", node.Val)
		InfixOrder(node.Right)
	}
}
//②、前序遍历:根节点——》左子树——》右子树
func PreOrder(node *linked_list.BinaryNode) {
	if node != nil {
		fmt.Printf("前序遍历:%v\n", node.Val)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
//③、后序遍历:左子树——》右子树——》根节点
func SufOrder(node *linked_list.BinaryNode) {
	if node != nil {
		SufOrder(node.Left)
		SufOrder(node.Right)
		fmt.Printf("后序遍历:%v\n", node.Val)
	}
}