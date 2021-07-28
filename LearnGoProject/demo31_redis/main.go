package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)
import "context"

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
	DB: 0,
})

/*
redis 学习内容参考  http://redisdoc.com/topic/cluster-spec.html
*/
func TestHash()  {

	//sliceKeys, err := rdb.HKeys(ctx, "stu2").Result()
	//if err == nil {
	//	for _, key := range sliceKeys {
	//		rdb.HDel(ctx, "stu2", key)
	//	}
	//}

	//设置给定key的field和值. 如给定的key不存在, 则创建. 内部必须键值(field-value)对应
	//   - HSet("myhash", "key1", "value1", "key2", "value2")
	//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
	//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
	cmd := rdb.HSet(ctx, "stu1","name", "张三", "age", 12)
	fmt.Println(cmd)

	//当指定的field不存在时, 赋值
	fmt.Printf("HSetNX:%v\n", rdb.HSetNX(ctx, "stu1", "age", "女"))

	//获取指定键是否存在
	fmt.Printf("获取指定键的field是否存在 %v\n", rdb.HExists(ctx, "stu1", "age"))
	fmt.Printf("获取指定键的field是否存在 %v\n", rdb.HExists(ctx, "stu2", "age"))

	//
	rdb.HMSet(ctx, "stu2", "Name", "李四", "age", 22, "Name", "王五")
	//删除指定的field下面的键值
	//rdb.HDel(ctx, "stu2", "Name")
	//不指定field, 错误.     必须带入指定的field
	fmt.Printf("hdel 不指定field %v\n", rdb.HDel(ctx, "stu2"))

	//获取指定key的所有field-value
	fmt.Printf("hgetall: %v\n", rdb.HGetAll(ctx, "stu2"))

	//获取key下面的字段数量
	fmt.Printf("hlen:%v\n", rdb.HLen(ctx, "stu2"))

	//hstrlen  没有对应的方法.   用于获取指定key的某一个field的值的长度
	//获取指定key下面的所有field和values
	fmt.Printf("hkeys:%v, hvals:%v\n", rdb.HKeys(ctx, "stu2"), rdb.HVals(ctx, "stu2"))

	//设置某一个field的增量  + incr_value
	rdb.HIncrBy(ctx, "stu2", "age", 10)
	fmt.Printf("hincrby age:%v\n", rdb.HGet(ctx, "stu2", "age"))

	ageMap := make(map[string]interface{})
	ageMap["age1"] = 01
	ageMap["age2"] = 02
	ageMap["age3"] = 03
	ageMap["age4"] = 04
	ageMap["age5"] = 05
	ageMap["age6"] = 06
	ageMap["age7"] = 07
	ageMap["age8"] = 8
	ageMap["age9"] = 9
	rdb.HMSet(ctx, "stu2", ageMap)

	//查找key对应的field 返回找到的字段和对应的游标 在量大的时候, 没有hkeys耗时
	//cursor设置为0时, 将会开启新一次的迭代. 迭代后将会返回一个新的cursor, 用于用户下一次的迭代
	cmdScan := rdb.HScan(ctx, "stu2", 3, "a*", 1)
	fmt.Println(cmdScan)
	fmt.Println(cmdScan.Result())
}
func main()  {
	TestHash()
	//TestList()
	//TestSet()
	//TestZSet()
	//TestHyperLogLog()
	//TestGEO()
	//TestBitMap()
	//TestDatabase()
	//TestPublishAndSubcribe()
	//TestRedisSystemCmd()
	//TestBuildinDebug()
}


