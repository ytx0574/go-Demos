package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

/*
zset  有序的集合,  每一个元素都包含score和member, 内部根据score来排序. 相同score的member的排序基于member的字典序来排序(lexicographical ordering)
*/
func TestZSet() {

	var sliceLen = 10
	var slice []*redis.Z = make([]*redis.Z, 10)
	for i := 0; i < sliceLen; i++ {
		slice[i] = &redis.Z{Score: float64(i), Member: string('a' + i) }
		fmt.Printf("score:%v, member:%v\n", slice[i].Score, slice[i].Member)
	}

	defaultZset := "zset1"

	//todo:根据带入的score大小, 移除所有zset中的元素
	//rdb.ZRemRangeByScore(ctx, defaultZset, "0", "10")
	//todo:根据member字典序来移除zset元素
	//rdb.ZRemRangeByLex(ctx, defaultZset, "-", "+")
	//todo:根据member排名来移除zset元素
	rdb.ZRemRangeByRank(ctx, defaultZset, 0, 10)
	//todo:获取zset中所有的元素, 包含返回score
	//z, err := rdb.ZRangeWithScores(ctx, defaultZset, 0, -1).Result()
	//for _, v := range z {
	//	//todo:移除zset中的元素.  带入的参数为member
	//	fmt.Println(rdb.ZRem(ctx, defaultZset, v.Member))
	//
	//}
	printZSetValues(defaultZset)


	rdb.ZAdd(ctx, defaultZset, &redis.Z{Score: 1, Member: "aa"}, &redis.Z{Score: 2, Member: "bb"}, slice[0], slice[2], slice[9], slice[1])
	printZSetValues(defaultZset)

	//todo://返回zset中的元素个数
	count, err := rdb.ZCard(ctx, defaultZset).Result()
	fmt.Printf("zcard %v, err:%v\n", count, err)
	//todo:返回指定的member的score
	fmt.Printf("zset score:%v\n", rdb.ZScore(ctx, defaultZset, "bb"))
	//todo:返回指定score范围的个数
	fmt.Printf("zset count:%v\n", rdb.ZCount(ctx, defaultZset, "1", "2"))

	//todo:按score值递减来查询
	fmt.Printf("zset revrange:%v\n", rdb.ZRevRangeWithScores(ctx, defaultZset, 0, -1))

	//todo:根据score范围来获取元素, 带入的参数为zrangeby指针
	zrangeby := &redis.ZRangeBy{Min: "0", Max: "3", Offset: 0, Count: 1}
	fmt.Printf("zset range by score:%v\n", rdb.ZRangeByScoreWithScores(ctx, defaultZset, zrangeby))

	//todo:获取指定member的排序位置
	fmt.Printf("zset zrank:%v\n", rdb.ZRank(ctx, defaultZset, "j"))
	fmt.Printf("zset zrevrank:%v\n", rdb.ZRevRank(ctx, defaultZset, "j"))

	//todo:返回指定的member的字典序的元素 此处的min max指定值特殊处理
	//合法的 min 和 max 参数必须包含 ( 或者 [ ， 其中 ( 表示开区间（指定的值不会被包含在范围之内）， 而 [ 则表示闭区间（指定的值会被包含在范围之内）。

	//特殊值 + 和 - 在 min 参数以及 max 参数中具有特殊的意义， 其中 + 表示正无限， 而 - 表示负无限。
	//因此， 向一个所有成员的分值都相同的有序集合发送命令 ZRANGEBYLEX <zset> - + ， 命令将返回有序集合中的所有元素。
	zrangeby = &redis.ZRangeBy{Min: "-", Max: "[bb", Offset: 0, Count: 10}
	zrangeby = &redis.ZRangeBy{Min: "[c", Max: "[bb", Offset: 0, Count: 10}
	zrangeby = &redis.ZRangeBy{Min: "(b", Max: "+", Offset: 0, Count: 10}
	fmt.Printf("zset range by lex:%v\n", rdb.ZRangeByLex(ctx, defaultZset, zrangeby))

	//todo:根据member字典序获取指定范围的元素个数
	fmt.Printf("zset zlexcount:%v\n", rdb.ZLexCount(ctx, defaultZset, "[a", "[b"))

	//todo://迭代获取指定正则匹配的member, 来查询
	fmt.Printf("zset scan:%v\n", rdb.ZScan(ctx, defaultZset, 0, "[a-b]", 10))

	//todo://对指定的member的score增量
	rdb.ZIncrBy(ctx, defaultZset, 30, "c")
	printZSetValues(defaultZset)

	defaultZset2 := "zset2"
	defaultZset3 := "zset3"
	rdb.ZRemRangeByRank(ctx, defaultZset2, 0, 1000000)
	rdb.ZRemRangeByRank(ctx, defaultZset3, 0, 1000000)
	rdb.ZAdd(ctx, defaultZset2, &redis.Z{Member: "李四", Score: 10000}, &redis.Z{Member: "赵六", Score: 8000}, &redis.Z{Member: "aston", Score: 9000})
	rdb.ZAdd(ctx, defaultZset3, &redis.Z{Member: "john", Score: 7000}, &redis.Z{Member: "aston", Score: 8200})
	printZSetValues(defaultZset2)
	printZSetValues(defaultZset3)

	//todo:从多个zset获取并集, 并归类到一个新的zset.  weight为乘法因子, 可控制合并进入的集合的score值. 这里单独对zset3的score*3
	//Aggregate 表示多个zset中相同member合并后, Score的取值方式.  SUM为所有的score和. MIN取最小 MAX取最大. 最终取值时是与乘法因子相乘后再比较
	zstore := redis.ZStore{
		Keys: []string{defaultZset2, defaultZset3},
		Weights: []float64{1.2, 3.0},
		Aggregate: "MIN",
	}
	defaultZsetUnion := "zset_union"
	zunionCmd := rdb.ZUnionStore(ctx, defaultZsetUnion, &zstore)
	fmt.Printf("zunioncmd:%v\n", zunionCmd)
	printZSetValues(defaultZsetUnion)

	defaultZsetInter := "zset_inter"
	zstore.Weights[0] = 1
	zstore.Weights[1] = 2
	rdb.ZInterStore(ctx, defaultZsetInter, &zstore)
	printZSetValues(defaultZsetInter)
}

func printZSetValues(key string)  {
	fmt.Printf("Zet All Values:%v\n", rdb.ZRangeWithScores(ctx, key, 0, -1))
}