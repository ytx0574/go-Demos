package main

import (
"fmt"
"sort"
"strconv"
)

//map和slice一样  都是引用类型  而引用类型要使用, 必须先初始化
//map不需要像数组 slice一样指定大小.  内容可无限
//map的key是无序
//map的数据删除使用delete. 如果删除全部. 则使用make重新分配一块空间. 旧的数据自动被gc回收

//map使用细节:
//1.map为引用类型, 遵循引用传递机制. 任何地方修改, 都会改变原来的值
//2.map不需要像数组一样指定大小. 可以无限添加数据
func main() {
	//mpa的初始化
	//第一种, 先声明再初始化. 可以理解为new
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["11"] = "11"
	fmt.Println(map1)

	//第二种 直接初始化
	map2 := make(map[int]string)
	map2[1] = "111"
	map2[1] = "333"
	map2[2] = "222"
	fmt.Printf("map2 = %v, map2 len = %v\n", map2, len(map2))

	//第三种  直接初始化
	map3 := map[float64]string {1:"111", 2: "222"}
	map4 := map[float64]string {
		1:"111",
		2: "222",   //换行初始化时, 必须补逗号
	}
	fmt.Println("map3 = ", map3, "map4 = ", map4)
	var num  = 11
	for i := 0; i < num; i++ {
		map2[i] = strconv.Itoa(i * 10000)
	}
	fmt.Printf("map2 = %v, map2 len = %v\n", map2, len(map2))

	var map5 = make(map[string]map[int]string)
	map5["第一个同学"] = make(map[int]string)
	map5["第一个同学"][1] = "11"
	map5["第一个同学"][2] = "22"

	map5["第二个同学"] = make(map[int]string)
	map5["第二个同学"][2] = "22"
	fmt.Printf("map5 = %v\n", map5)

	delete(map5, "aaa")
	fmt.Printf("map5 = %v\n", map5)
	delete(map5, "第一个同学")
	fmt.Printf("map5 = %v\n", map5)


	map6 := map[string]string {
		"3" : "33",
		"33" : "33",
		"1" : "11",
		"11" : "11",
		"2" : "22",
		"22" : "22",
	}
	fmt.Printf("map6 = %v\n", map6)
	var keys []int = make([]int, len(map6))
	for key, _ := range  map6 {
		num, _ := strconv.ParseInt(key, 10, 0)
		keys = append(keys, int(num))
	}

	sort.Ints(keys) //排序keys

	for _, num := range  keys {
		key := strconv.Itoa(num)
		fmt.Printf("顺序输出map6 key:%v, value:%v\n", key, map6[key])
	}


}
