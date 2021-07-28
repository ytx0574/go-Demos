package main

import (
	"fmt"
	"time"
)

/*
list 内部由链表来实现, 首尾填充数据很快, 数据量大的时候, 遍历效果一般
*/
func TestList() {
	//先移除所有值
	var defaultListKey = "list1"

	list1Values, err := rdb.LRange(ctx, defaultListKey, 0, -1).Result()
	if err == nil {
		 for _, v := range list1Values {
		 	//删除list中指定的值
		 	rdb.LRem(ctx, defaultListKey, 0, v)
		 }
	}


	//list的默认实现是一个链表.  内部展示数据为 先入的数据在右边. 如 push a, b, c, 内部展示为 c b a
	//从左往右推入一个值到list. 当指定key值不存在时, 会针对该key创建一个list
	rdb.LPush(ctx, defaultListKey, "a", "b", "c")
	printAllListValues(defaultListKey)

	rdb.RPop(ctx, defaultListKey)
	printAllListValues(defaultListKey)

	rdb.LPushX(ctx, defaultListKey, "d", "e")
	printAllListValues(defaultListKey)

	//从右推一个值到列表. 当指定的key不存在时, 不做任何操作
	rdb.RPushX(ctx, defaultListKey, "ff", "gg", "gg")
	printAllListValues(defaultListKey)


	//删除list中和指定value相同的值. count>0 从左往右开始, count<0 从右往左开始.
	//删除指定count的绝对值的数量
	rdb.LRem(ctx, defaultListKey, 1, "gg")
	printAllListValues(defaultListKey)

	fmt.Printf("获取指定list的长度:%v\n", rdb.LLen(ctx, defaultListKey))

	fmt.Printf("获取指定list的指定位置的值:%v\n", rdb.LIndex(ctx, defaultListKey, 3))

	//從左往右找, 在找到的第一個值的前後插入值, 找不到則不插入, 返回-1
	rdb.LPush(ctx, defaultListKey, "大", "小")
	fmt.Printf("插入:%v\n", rdb.LInsertAfter(ctx, defaultListKey, "c", "大"))
	fmt.Printf("插入:%v\n", rdb.LInsertBefore(ctx, defaultListKey, "大", "小"))
	fmt.Printf("插入:%v\n", rdb.LInsertBefore(ctx, defaultListKey, "大", "小"))
	fmt.Printf("插入:%v\n", rdb.LInsertBefore(ctx, defaultListKey, " ", "小"))
	printAllListValues(defaultListKey)

	//替換指定位置的值
	rdb.LSet(ctx, defaultListKey, 2, "小小")
	printAllListValues(defaultListKey)

	//保留指定位置的值
	rdb.LTrim(ctx, defaultListKey, 1, -1)
	printAllListValues(defaultListKey)

	//blpop brpop brpoplpush 是lpop rpop rpoplpush的阻塞版本, 当list中没有数据时, 会阻塞, 直到超时才会继续执行
	fmt.Printf("blpop: %v\n", rdb.BLPop(ctx, time.Second * 1, defaultListKey))
	printAllListValues(defaultListKey)
	list1Values, err = rdb.LRange(ctx, defaultListKey, 0, -1).Result()
	//清空默认的list1的数据
	for i, v := range  list1Values {
		fmt.Printf("list1 index:%v, value:%v\n", i, v)
		rdb.RPop(ctx, defaultListKey)
	}
	//操作一个空list或一个不存在的list 会返回redis.nil
	fmt.Printf("blpop: %v\n", rdb.BLPop(ctx, time.Second * 1, defaultListKey))
	fmt.Printf("brpoplpush: %v\n", rdb.BRPopLPush(ctx, defaultListKey, defaultListKey, time.Second * 1))
	printAllListValues(defaultListKey)


	rdb.LPush(ctx, defaultListKey, "a", "b", "c")
	defaultListKey2 := "list2"
	rdb.RPush(ctx, defaultListKey2, "11", "22", 33)

	printAllListValues(defaultListKey)
	printAllListValues(defaultListKey2)

	//把source list的最右边一个移出, 换到destination list的左边
	rdb.RPopLPush(ctx, defaultListKey2, defaultListKey)
	printAllListValues(defaultListKey)
	printAllListValues(defaultListKey2)

	//操作同一个list, 则最右一个变换到最左, 形成了一个旋转(rotation)
	rdb.RPopLPush(ctx, defaultListKey, defaultListKey)
	printAllListValues(defaultListKey)


	rdb.BRPopLPush(ctx, defaultListKey, defaultListKey, time.Second * 1)
	printAllListValues(defaultListKey)
}

func printAllListValues(key string) {
	//lrange 獲取指定範圍的值, 負數從右往左開始.
	fmt.Printf("获取指定list的值:%v\n", rdb.LRange(ctx, key, 0, -1))
}
