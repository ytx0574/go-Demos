package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)


/*
此处的效果要结合redis server来理解  因为server一直开着, 所以第一运行后, 后面会获取之前的数据
*/
//func Ctx() context.Context {
//	return ctx
//}
//func Rdb() *redis.Client {
//	return rdb
//}

func TestString()  {
	//常规的键值保存
	var cmd *redis.StatusCmd = rdb.Set(ctx, "key", "value", 0)
	fmt.Println(cmd)
	cmdGet := rdb.Get(ctx, "key")
	fmt.Println(cmdGet)

	cmd1 := rdb.Set(ctx, "大哥", 111, time.Second * 5)
	fmt.Println(cmd1)
	fmt.Println(cmd1.Args())

	cmd1Get := rdb.Get(ctx, "大哥")
	fmt.Println(cmd1Get)
	fmt.Println(cmd1Get.Args())


	//KeepTTL 使用之前的过期时间继续设置字段. 比如之前设置的过期时间为5s, 3s后再次更改, 那么过期时间又从3s开始
	cmd1 = rdb.Set(ctx, "大哥", 222, redis.KeepTTL)
	fmt.Println(cmd1)
	cmd1Get = rdb.Get(ctx, "大哥")
	_, err := cmd1Get.Result()
	if err != redis.Nil {
		fmt.Println(cmd1Get.Args()) //获取入参
 		fmt.Println(cmd1Get.Result()) //获取结果
	}

	//time.Sleep(time.Second * 2) //延迟后 再次获取有效时间

	//获取字段的过期时间
	//此处的非常规操作, 一般不好理解.   redis.Client继承自combale这个Func类型, 内部自行设置了继承变量的值. 调用时通过变量来调用
	var cmdDuration *redis.DurationCmd = rdb.TTL(ctx, "大哥")
	cmdDuration.Val()
	fmt.Printf("获取key的有效时间 %v\n", cmdDuration)


	rdb.SetNX(ctx, "name", "dating", 0)
	fmt.Printf("获取setnx的值:%v\n", rdb.Get(ctx, "name"))

	// NX当值不存在时, 才赋值, 否则无法覆盖
	rdb.SetNX(ctx, "name", "dating 11", 000000000000000000000)
	fmt.Printf("获取setnx的值:%v\n", rdb.Get(ctx, "name"))


	// XX当值存在时 才能赋值
	rdb.SetXX(ctx,"name", "name1 value", 0)
	fmt.Printf("获取setxx的值:%v\n", rdb.Get(ctx, "name"))

	//和set一样的效果, setex是一个原子操作, 它可以在同一时间完成设置值和过期时间两个操作
	//setex的单位必须是s  不可带入小于1s的数值
	rdb.SetEX(ctx, "name", "大宝贝", time.Millisecond * 1000)
	fmt.Printf("获取setex的值:%v\n", rdb.Get(ctx, "name"))
	//time.Sleep(time.Millisecond * 1001)
	fmt.Printf("获取setxx的值:%v\n", rdb.Get(ctx, "name"))

	//获取值的时候, 设置值. 并返回之前的值, 如果直接没有返回nil
	fmt.Printf("getset :%v\n", rdb.GetSet(ctx, "大哥", "大哥 new value"))
	fmt.Printf("getset :%v\n", rdb.GetSet(ctx, "大", "123456789"))


	//获取字符串值的长度
	fmt.Printf("获取字符串值的长度:%v\n", rdb.StrLen(ctx, "大"))
	//追加字符串值  对于不存在的字符串 等同于set
	rdb.Append(ctx, "大", "appending string ~~~~")
	//fmt.Printf("getset :%v\n", rdb.GetSet(ctx, "大", "大  new value"))
	fmt.Printf("获取字符串值的长度:%v\n", rdb.StrLen(ctx, "大"))

	//设置字符串的值 包含偏移量, 不存在的key, 直接覆盖值(不够的长度用零字节展示"\x00", 打印时看不出来, 实际的长度还是你偏移后的长度)
	rdb.SetRange(ctx, "大", 40, "\"offset stirng\"")
	rdb.SetRange(ctx, "大", 20, "\"(20)\"")
	fmt.Printf("获取字符串值的值:%v\n", rdb.Get(ctx, "大"))
	fmt.Printf("获取字符串值的长度:%v\n", rdb.StrLen(ctx, "大"))

	rdb.Append(ctx, "大", "Append Again大")
	fmt.Printf("获取字符串值的值:%v\n", rdb.Get(ctx, "大"))
	fmt.Printf("获取字符串值的长度:%v\n", rdb.StrLen(ctx, "大"))

	rdb.SetRange(ctx, "RangeKey", 30, "aa test")
	fmt.Printf("setrange设置不存在的key, 获取结果:%v len:%v\n", rdb.Get(ctx,"RangeKey"), rdb.StrLen(ctx, "RangeKey"))

	//获取指定key的值的区间值  start end分别代指下标. 0, 10前零到11个字符   -1 最后一个字符 -2倒数第二个字符 超过下标的部分会被自动忽略
	getrange := rdb.GetRange(ctx, "大", 0, 5)
	if _, err := getrange.Result(); err == redis.Nil {
		fmt.Printf("getrange 无效\n")
	}
	fmt.Printf("getrange :%v\n", getrange)

	//为设置的key的值+1 如果key的值不能被解释为数字, 此命令会返回错误
	fmt.Printf("为不能解释为整数的值+1 %v\n", rdb.Incr(ctx, "大"))
	rdb.Set(ctx, "age", 10, 0)
	rdb.Incr(ctx, "age")
	fmt.Printf("+1后的结果:%v\n", rdb.Get(ctx, "age"))

	//为指定key的值升值增量, 指定增加的值. 如果key不存在, 那么key会先被初始化为0, 然后再添加增量
	rdb.IncrBy(ctx, "age", 10)
	fmt.Printf("+1后的结果:%v\n", rdb.Get(ctx, "age"))

	//自减-
	rdb.Decr(ctx, "age")
	//自减指定减量
	rdb.DecrBy(ctx, "age", 20)

	//多key赋值. 如果key value 不对应, 返回错误 所有的key都无法设置
	cmdmset := rdb.MSet(ctx, "name1", "aston", "age1", "12", "gender1", "女")
	fmt.Printf("mset:%v\n", cmdmset)
	fmt.Printf("name1:%v, age1:%v, gender1:%v\n", rdb.Get(ctx, "name1"), rdb.Get(ctx, "age1"), rdb.Get(ctx, "gender1"))

	rdb.MSet(ctx, "name", "aston", "age", "88", "gender", "♂")
	fmt.Printf("name:%v, age:%v, gender:%v\n", rdb.Get(ctx, "name"), rdb.Get(ctx, "age"), rdb.Get(ctx, "gender"))

	//mset mget 同时设置多个值
	//mset支持指定格式的map  map[string]interface{}  其他格式无效
	//支持指定格式的数组 如 []string{key, value, key1, value1}
	myMap := make(map[string]interface{})
	myMap["MName"] = "aa"
	myMap["MAge"] = "15"
	fmt.Printf("mset map:%v\n", rdb.MSet(ctx, myMap))
	fmt.Printf("mget:%v\n", rdb.MGet(ctx, "MName", "MAge", "MGender"))
}