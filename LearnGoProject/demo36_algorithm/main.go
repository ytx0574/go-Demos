package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*
选择排序
从已知列里面递归循环, 找出每一次循环的最大值, 然后与本次循环的第一个值进行交换
*/
func SelectSort(ay []int) {

	for i := 0; i < len(ay) - 1; i++ {

		max := ay[i]
		maxIdx := i
		//假定第一个就是最大的, 从i+1开始, 找到后面数值中最大的
		for j := i + 1; j < len(ay); j++ {
			if max < ay[j] {
				max = ay[j]
				maxIdx = j
			}
		}
		//如
		if maxIdx != i {
			ay[i], ay[maxIdx] = ay[maxIdx], ay[i]
		}
	}
}

/*
插入排序
把数据分为两个表, 一个有序 一个无序.  开始时, 有序表只有一个元素, 排序时每次从无序表中取出第一个元素, 把它与有序表进行比较, 然后插入到合适位置
重点在于 脑海中形成逻辑, 每一次循环, 会移动一个数, 从前往后移.(不用担心移动后有两个相同的数, 因为再次移动会把更前面的数拿过来覆盖) 比如
[1, 2, 3]
[1, 1] -> [2, 1]
[2, 1, 1] -> [2, 2, 1] -> [3, 2, 1]
*/
func InsertSort(ay []int) {
	//假定第一个数为有序, 那么从第二个数开始循环
	//for i := 1; i < len(ay); i++ {
	//	insertValue := ay[i]
	//	insertIndex := i - 1  //每次循环开始前, 被插入的下标为其-1
	//
	//	//假定把第二个数和前面只有一个数的有序表进行比对. 那么插入的比对就是 1 - 1
	//	for insertIndex >= 0 && ay[insertIndex] < insertValue {
	//		ay[insertIndex + 1] = ay[insertIndex] //如果当前值小于被插入的值, 那么当然值就往后移一位
	//		insertIndex--
	//	}
	//
	//	//找出最后的下标, 是否和插入值相同, 不相同就替换掉插入值
	//	if ay[insertIndex + 1] != insertValue {
	//		ay[insertIndex + 1] = insertValue
	//	}
	//}

	for i := 1; i < len(ay); i++ {
		insertValue := ay[i]
		insertIndex := i

		//取前面一个数与之比较, 小于就继续把数值往后移
		for insertIndex > 0 && ay[insertIndex - 1] < insertValue {
			ay[insertIndex] = ay[insertIndex - 1]
			insertIndex--
		}

		//如果两个不是一个位置 则交换两者的值
		if i != insertIndex {
			ay[insertIndex] = insertValue
		}

	}
}

/*
快速排序
选择第一个数为基准值, 从它后面开始遍历, 只要遇到比它小的值, 就把它往前挪. 最后再把基准值和最后一个数进行交换. 递归执行
*/
func QuickSort(left, right int, ay[]int) {

	swap := func(i, j int, ay[]int) {
		if i == j {  //两者相等时, 不可用异或交换法. (临时变量交换的可以的)
			return
		}
		ay[i] = ay[i] ^ ay[j]
		ay[j] = ay[j] ^ ay[i]
		ay[i] = ay[i] ^ ay[j]
	}

	partition := func(left, right int, ay[]int) int {
		pivot := left //以第一个数为基准值
		idx := pivot + 1
		//从基准值后面开始遍历其他值, 把小于的搬到前面
		for i := idx; i <= right; i++ {
			if ay[i] < ay[pivot] { //小于基准值, 则和前面的数值进行交换
				swap(i, idx, ay)
				idx++
			}
		}
		//移动位置后, 基准值的坐标就是-1. 因为上面多加
		idx--
		swap(pivot, idx, ay)
		return idx
	}

	if left < right {
		idx := partition(left, right, ay)
		QuickSort(left, idx - 1, ay)
		QuickSort(idx + 1, right, ay)
	}
}
/*
快速排序
选择第一个数为基准值, 从右往左找一个小于基准值的下标, 从左往右找一个大于基准值的下标. 在满足条件的情况下进行值交换. 最后得到的就是左边的小, 后边的大. 再把基准值插到中间
*/
func QuickSort2(left, right int, ay []int) {
	if left < right {
		pivot := ay[left] //以第一个数值为基准值
		i := left
		j := right
		for i < j {
			//从右往左找, 遇到小于基准值, 跳出循环.
			//顺序很重要，要先从右边开始找 (因为基准数是第一个, 那么必须先从它的对立面开始数)
			//原因参考 https://blog.csdn.net/lkp1603645756/article/details/85008715
			for ay[j] >= pivot && i < j {
				j--
			}
			//从左往右找, 找到大于基准值的下标  (跳出循环时, 就意味着有可能找到大于基准值的位置)
			for ay[i] <= pivot && i < j {
				i++
			}

			//如前后者找到, 则进行值交换
			if i < j {
				ay[i] = ay[i] ^ ay[j]
				ay[j] = ay[i] ^ ay[j]
				ay[i] = ay[i] ^ ay[j]
			}
		}

		//因为i小于基准值, 那么再把i的值于基准值进行交换
		ay[left] = ay[i]
		ay[i] = pivot

		QuickSort2(left, i - 1, ay)
		QuickSort2(i + 1, right, ay)
	}
}

func QuickSort3(left int, right int, array []int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left + right) / 2]
	temp := 0
	//for 循环的目标是将比 pivot 小的数放到 左边 // 比 pivot 大的数放到 右边
	for ; l < r; {
		//从 pivot 的左边找到大于等于 pivot 的值
		for ; array[l] < pivot; {
			l++
		}
		//从 pivot 的右边边找到小于等于 pivot 的值
		for ; array[r] > pivot; {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r {
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l]== pivot {
			r--
		}
		if array[r]== pivot {
			l++
		}
	}
	// 如果
	if l == r {
		l++
		r--
	}
	// 向左递归
	if left < r {
		QuickSort3(left, r, array)
	}
	// 向右递归 1== r, 再移动下
	if right > l {
		QuickSort3(l, right, array)
	}
}

//求阶乘
//乘方转乘法
// x ^ y;
//a = (x * x)  b = (y / 2);
//x ^ y = a ^ b
//b为奇数时, 额外乘一个a
func pow(x, y int) int {
	if x == 0 || x == 1 {
		return x
	}

	if y > 1 {
		b := y / 2
		a := x * x
		if y % 2 == 1 {
			return pow(a, b) * x
		}else {
			return pow(a, b)
		}
	}else if (y == 0) {
		return 1
	}else {
		return x
	}
}


//https://zhuanlan.zhihu.com/p/37468694
//背包问题  从已知的数值中组合出固定数值
/*
一、如果在这个过程的任何时刻，选择的数据项的总和符合目标重量，那么工作便完成了。
二、从选择的第一个数据项开始，剩余的数据项的加和必须符合背包的目标重量减去第一个数据项的重量，这是一个新的目标重量。
三、逐个的试每种剩余数据项组合的可能性，但是注意不要去试所有的组合，因为只要数据项的和大于目标重量的时候，就停止添加数据。
四、如果没有合适的组合，放弃第一个数据项，并且从第二个数据项开始再重复一遍整个过程。
五、继续从第三个数据项开始，如此下去直到你已经试验了所有的组合，这时才知道有没有解决方案。
*/
type Knapsack struct {
	weights []int
	selects []bool
}

func (this *Knapsack)knapsack(total, index int) {
	if total < 0 || (total > 0 && index >= len(this.weights)) {
		return
	}
	if total == 0 {
		for i := 0; i < index; i++ {
			if this.selects[i] {
				fmt.Printf("%v ", this.weights[i])
			}
		}
		fmt.Println()
		return
	}
	this.selects[index] = true
	this.knapsack(total - this.weights[index], index + 1)
	this.selects[index] = false
	this.knapsack(total, index + 1)
}


/*
组合选择
*/
type Combination struct {
	persons []string
	selects []bool
}

func (this *Combination)showTeams(count int) {
	this.combination(count, 0)
}

//func (this *Combination)combination(teamnumber, index int) {
//	if teamnumber == 0 {
//		for i := 0; i < len(this.selects); i++ {
//			if this.selects[i] {
//				fmt.Printf("%v ", this.persons[i])
//			}
//		}
//		fmt.Println()
//	}
//
//	if index >= len(this.persons) {
//		return
//	}
//
//	this.selects[index] = true
//	this.combination(teamnumber - 1, index + 1)
//	this.selects[index] = false
//	this.combination(teamnumber, index + 1)
//}

func main() {
	const maxCount = 6
	ay := [maxCount]int{33, 50, 55, 22, 3, 99}

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 6; i < maxCount; i++ {
		ay[i] = rand.Intn(99)
	}

	ayBackup := ay
	t := time.Now()
	fmt.Printf("排序前的数组:%v\n", ay)
	SelectSort(ay[:])
	fmt.Printf("选择排序后的数组:%v  %v\n", ay, time.Now().Sub(t))


	ay = ayBackup
	t = time.Now()
	fmt.Printf("排序前的数组:%v\n", ay)
	InsertSort(ay[:])
	fmt.Printf("插入排序后的数组:%v  %v\n", ay, time.Now().Sub(t))

	ay = ayBackup
	t = time.Now()
	fmt.Printf("排序前的数组:%v\n", ay)
	QuickSort(0, len(ay) - 1, ay[:])
	fmt.Printf("快速排序后的数组:%v  %v\n", ay, time.Now().Sub(t))

	ay = ayBackup
	t = time.Now()
	fmt.Printf("排序前的数组:%v\n", ay)
	QuickSort2(0, len(ay) - 1, ay[:])
	fmt.Printf("快速排序后的数组:%v  %v\n", ay, time.Now().Sub(t))


	ay = ayBackup

	sort.Slice(ay[:], func(i, j int) bool {
		return ay[i] < ay[j]
	})
	fmt.Println(ay)


	fmt.Println(pow(2, 6))

	copy(ay[:], []int{7, 11, 6, 5, 9, 0})
	fmt.Println(ay)
	(&Knapsack{
		weights: ay[:],
		selects: make([]bool, len(ay)),
	}).knapsack(16, 0)


	//(&Combination{
	//	persons: []string{"A", "B", "C", "D", "E"},
	//	selects: make([]bool, 5),
	//}).showTeams(4)
}