package main

import "fmt"

/*
set  不可出现相同的元素
*/
func TestSet(){
	defaultSetName := "set"
	//TODO:向set中添加元素, 返回添加元素的数量
	fmt.Printf("set add:%v\n", rdb.SAdd(ctx, defaultSetName, 1, 2, 3, 4, 5, 3, "a", "b", "oo"))
	printAllMembers(defaultSetName)

	//TODO:获取某个值是否存在set中
	fmt.Printf("sismember:%v\n", rdb.SIsMember(ctx, defaultSetName, 4))
	fmt.Printf("sismember:%v\n", rdb.SIsMember(ctx, defaultSetName, 41))

	//todo:随机移除set中的一个元素
	printAllMembers(defaultSetName)
	rdb.SPop(ctx, defaultSetName)
	printAllMembers(defaultSetName)

	//todo:随机获取set中的一个元素, 不从set移除
	fmt.Printf("srandmember:%v\n", rdb.SRandMember(ctx, defaultSetName))
	printAllMembers(defaultSetName)

	//todo:随机从set中获取N个元素, n > set.count, 返回count个元素
	fmt.Printf("srandmember:%v\n", rdb.SRandMemberN(ctx, defaultSetName, 2))
	printAllMembers(defaultSetName)

	//todo:移除set中指定值
	rdb.SRem(ctx, defaultSetName, 5)
	rdb.SRem(ctx, defaultSetName, 44)
	printAllMembers(defaultSetName)

	//todo:根据cursor迭代查找根据正则查找set中的值, 下次迭代可从返回的cursor开始迭代
	fmt.Printf("set scan:%v\n", rdb.SScan(ctx, defaultSetName, 0, "[0-9]", 10))

	//todo:获取set的个数
	fmt.Printf("set count:%v\n", rdb.SCard(ctx, defaultSetName))

	defaultSetName2 := "set2"
	rdb.SAdd(ctx, defaultSetName2, "x", "y", "z")

	//todo:从source set移动指定的元素到destination set
	rdb.SMove(ctx, defaultSetName, defaultSetName2,  1)
	printAllMembers(defaultSetName)
	printAllMembers(defaultSetName2)


	//todo://获取两个set中相同的元素 交集
	fmt.Printf("set inter: %v\n", rdb.SInter(ctx, defaultSetName, defaultSetName2))
	defaultSetName3 := "set3"
	//todo:获取两个set中的相同元素到另一个集合   带store就是合并到另外一个set中, 不移除原set中的元素
	rdb.SInterStore(ctx, defaultSetName3, defaultSetName, defaultSetName2)
	printAllMembers(defaultSetName)
	printAllMembers(defaultSetName2)
	printAllMembers(defaultSetName3)

	//todo:获取两个set中的所有的元素 并集
	fmt.Printf("set union: %v\n", rdb.SUnion(ctx, defaultSetName, defaultSetName2))

	//todo:获取两个set中不同的元素 差集
	fmt.Printf("set diff: %v\n", rdb.SDiff(ctx, defaultSetName, defaultSetName2))

}

func printAllMembers(key string) {
	//todo://获取set中所有的元素
	fmt.Printf("All Members:%v\n", rdb.SMembers(ctx, key))
}
