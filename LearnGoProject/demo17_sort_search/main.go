package main

import "fmt"

func bubbleSort(arr *[]int) {

	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - 1 - i; j++ {
			//前者大于后者, 交换两个的位置
			value1 := (*arr)[j]
			value2 := (*arr)[j + 1]
			if value1 > value2 {
				(*arr)[j] = value2
				(*arr)[j + 1] = value1
			}
		}
	}

}

func binary_search1(arr *[]int, value int) (index int) {

	leftIndex := 0
	rightIndex := len(*arr) - 1

	var recursive func(arr *[]int, leftIndex, rightIndex, value int)
	//递归调用闭包自己
	recursive = func(arr *[]int, leftIndex, rightIndex, value int) {
		middleIndex := (leftIndex + rightIndex) / 2
		if (*arr)[middleIndex] > value {
			recursive(arr, leftIndex, middleIndex - 1, value)
		}else if (*arr)[middleIndex] < value {
			recursive(arr, middleIndex + 1, rightIndex, value)
		}else {
			index = middleIndex
		}
	}

	recursive(arr, leftIndex, rightIndex, value)

	defer func() {
		fmt.Printf("recursive: %v\n", recursive)
	}()
	return index
}

func binary_search(arr *[]int, leftIndex, rightIndex, value int) (index int) {
	if leftIndex > rightIndex {
		index = -1
	}

	middleIndex := (leftIndex + rightIndex) / 2
	if (*arr)[middleIndex] > value {
		//middleIndex本身大于值, 那么就从它下面的index开始
		return binary_search(arr, leftIndex, middleIndex - 1, value)
	} else if (*arr)[middleIndex] < value {
		//mindexIndex本身小于值, 那么从它上面的index开始
		return binary_search(arr, middleIndex + 1, rightIndex, value)
	} else {
		index = middleIndex
	}
	fmt.Printf("执行return...%v\n", index)
	return index //此处犯了一个错误, 递归调用是要反复执行的. 上面的递归调用没有添加return. 导致每一次递归都执行了return, 而每一次都是不一样的index
}


func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到 中间的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findVal { //说明我们要查找的数，应该在
		BinaryFind(arr, leftIndex, middle - 1, findVal)
	} else if (*arr)[middle] < findVal { //说明我们要查找的数，应该在 middel+1 --- rightIndex
		BinaryFind(arr, middle + 1, rightIndex, findVal)
	} else { //找到了
		fmt.Printf("找到了，下标为%v \n", middle)
	}
	fmt.Printf("执行return~...\n")
}

func main() {
	arr := []int{22, 3, 23, 42, 1, 23}
	fmt.Printf("排序前:%v\n", arr)
	bubbleSort(&arr)
	fmt.Printf("排序前:%v\n", arr)

	index := binary_search(&arr, 0, len(arr) - 1, 23)
	index1 := binary_search1(&arr, 23)
	fmt.Printf("binary_search index: %v\n", index)
	fmt.Printf("binary_search1 index1: %v\n", index1)

	//arr slice转arr1 原理, 从arr复制内容到arr1的切片. 因为arr1是引用类型, 所以对应的arr1数组内容也发生了变化
	var arr1 [6]int
	copy(arr1[:], arr[:])
	BinaryFind(&arr1, 0, len(arr1) - 1, 42)
}
