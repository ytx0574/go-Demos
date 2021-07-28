package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"strings"
)

/*

Redis HyperLogLog 是用来做基数统计的算法，HyperLogLog 的优点是，在输入元素的数量或者体积非常非常大时，计算基数所需的空间总是固定 的、并且是很小的。

在 Redis 里面，每个 HyperLogLog 键只需要花费 12 KB 内存，就可以计算接近 2^64 个不同元素的基 数。这和计算基数时，元素越多耗费内存就越多的集合形成鲜明对比。

可以理解为和set一样的效果, 但是又不会在数据量大的时候占用太多的内存, 对于数据量大的时候, 可以很方便的计算出总条数

比如: 统计注册ip数, 搜索词条数
*/
func TestHyperLogLog() {

	//todo:添加数据, 自动去重
	defaultHyperLogLog := "hyperloglog"
	rdb.PFAdd(ctx, defaultHyperLogLog, "1", "2", "大", "叫我姐佛教哦囧囧哦就哦哦哦")

	rdb.PFAdd(ctx, defaultHyperLogLog, "大")

	fmt.Printf("hyperLogLog count:%v\n", rdb.PFCount(ctx, defaultHyperLogLog))

	defaultHyperLogLog2 := "hyperloglog2"
	defaultHyperLogLog3 := "hyperloglo3"

	rdb.PFAdd(ctx, defaultHyperLogLog2, "2", "sojo", "jo2")
	rdb.PFAdd(ctx, defaultHyperLogLog3, "0", "沃", "32哦")

	//todo:合并多个hyperloglog到一个, 自动去重
	defaultHyperLogLogMerge := "hyperloglogmerge"
	rdb.PFMerge(ctx, defaultHyperLogLogMerge, defaultHyperLogLog, defaultHyperLogLog2, defaultHyperLogLog3)

	fmt.Printf("merge hyperloglog count:%v\n", rdb.PFCount(ctx, defaultHyperLogLogMerge))
}

func TestGEO() {
	defaultGeo := "geo"

	location1 := &redis.GeoLocation{Longitude: 11, Latitude: 22, Name: "location1"}
	rdb.GeoAdd(ctx, defaultGeo, location1)


	location2 := &redis.GeoLocation{Longitude: 33, Latitude: 44, Name: "location2"}
	rdb.GeoAdd(ctx, defaultGeo, location2)


	//todo:获取带入的member的经纬度
	fmt.Printf("location1:%v, location2:%v\n", location1, location2)
	cmdGeoPos := rdb.GeoPos(ctx, defaultGeo, "location1", "location2")
	fmt.Printf("获取指定元素的位置:%v\n", cmdGeoPos)
	geopos, _ := cmdGeoPos.Result()

	var geoPos []*redis.GeoPos = geopos
	for _, v := range  geoPos {
		fmt.Printf("Longitude:%v Latitude;%v\n", v.Longitude, v.Latitude)
	}

	//todo://获取两点的距离
	cmdGeoDist := rdb.GeoDist(ctx, defaultGeo, "location1", "location2", "km")
	fmt.Printf("获取两点距离:%v\n", cmdGeoDist)

	//todo:获取指定坐标的范围的位置数据
	var geoRadiusQuery *redis.GeoRadiusQuery = &redis.GeoRadiusQuery{Radius: 1000, Unit: "km", WithDist: true, WithCoord: true, Count: 10}
	fmt.Printf("获取指定范围的数据%v\n", rdb.GeoRadius(ctx, defaultGeo, 15, 30, geoRadiusQuery))

	//todo:根据内部的某一个member获取它指定范围内的数据, 返回数据包含它自己
	fmt.Printf("获取指定member指定范围的数据:%v\n", rdb.GeoRadiusByMember(ctx, defaultGeo, "location1", geoRadiusQuery))

	//todo:获取指定坐标的范文的位置数据, Store为一个指定的有序集合key.  geoRadiusQuery内部的参数Store or StoreUnit不能为空
	//todo: Store使用geohash作为zset的Score, 后者使用距离作为zset的Score
	geoRadiusQuery = &redis.GeoRadiusQuery{Radius: 1000, Unit: "km", Count: 10, Store: "zset3", StoreDist: ""}
	fmt.Printf("%v\n", rdb.GeoRadiusStore(ctx, defaultGeo, 15, 30, geoRadiusQuery))
	rdb.GeoRadiusByMemberStore(ctx, defaultGeo, "location2", geoRadiusQuery)

	//todo:获取指定member的geohash
	fmt.Printf("获取指定member的geohash:%v\n", rdb.GeoHash(ctx, defaultGeo, "location2"))

	//todo:提取储存的数据
	fmt.Println(rdb.ZRangeWithScores(ctx, "zset3", 0, -1))

}

/*
位图的值只能为0 - 1, 位图的offset必须>0, 小于2^32之间(512MB)
常见操作 用于统计登录用户 活跃用户 签到统计等
*/
var bitfieldkey = "bitfieldkey"
func TestBitMap() {
	defaultBitMap := "bitmap"

	//todo 设置获取指定位置的值
	fmt.Printf("设置位图指定位置的值:%v\n", rdb.SetBit(ctx, defaultBitMap, 11, 1))
	fmt.Printf("设置位图指定位置的值:%v\n", rdb.SetBit(ctx, defaultBitMap, 12, 0))

	fmt.Printf("获取位图指定位置的值:%v\n", rdb.GetBit(ctx, defaultBitMap, 11))
	fmt.Printf("获取位图指定位置的值:%v\n", rdb.GetBit(ctx, defaultBitMap, 12))

	//toto获取指定位置的中为1的数量
	var bitCount redis.BitCount = redis.BitCount{Start: 0, End: 100}
	fmt.Printf("获取位图中值为1的数量:%v\n", rdb.BitCount(ctx, defaultBitMap, &bitCount))

	//todo 获取bitmap中第一个位指定值的位置
	fmt.Printf("获取位图中第一个值为指定bit的offset:%v\n", rdb.BitPos(ctx, defaultBitMap, 1))



	defaultBitMap2 := "bitmap2"
	fmt.Printf("设置位图指定位置的值:%v\n", rdb.SetBit(ctx, defaultBitMap2, 11, 0))
	fmt.Printf("设置位图指定位置的值:%v\n", rdb.SetBit(ctx, defaultBitMap2, 12, 1))

	//todo://对指定的多个bitmap进行位操作, 并写入到一个新的bitmap
	//todo: and 逻辑与, or 逻辑或  xor逻辑异或  not 逻辑非
	defaultBitMapResult := "bitmap_result"
	fmt.Printf("获取位图数据 异或 xor:%v\n", rdb.BitOpXor(ctx, defaultBitMapResult, defaultBitMap, defaultBitMap2))
	fmt.Printf("获取位图数据 逻辑与 and:%v\n", rdb.BitOpAnd(ctx, defaultBitMapResult, defaultBitMap, defaultBitMap2))
	fmt.Printf("获取位图数据 逻辑或 or:%v\n", rdb.BitOpOr(ctx, defaultBitMapResult, defaultBitMap, defaultBitMap2))
	fmt.Printf("获取位图数据 逻辑非 not:%v\n", rdb.BitOpNot(ctx, defaultBitMapResult, defaultBitMap))
	fmt.Printf("获取位图指定位置的值:%v\n", rdb.GetBit(ctx, defaultBitMapResult, 11))
	fmt.Printf("获取位图指定位置的值:%v\n", rdb.GetBit(ctx, defaultBitMapResult, 12))



	//todo:bitfield使用参考:  http://redisdoc.com/bitmap/bitfield.html
	//1100100  [110001 110000 110000]
	fmt.Printf("%b\n", []byte("109"))
	rdb.Set(ctx, bitfieldkey, "109", 0)

	//todo:获取指定位的值
	fmt.Println(rdb.BitField(ctx, bitfieldkey, "get", "u8", 0))
	fmt.Println(rdb.BitField(ctx, bitfieldkey, "get", "u8", 16))
	//todo:将字符串看做二进制数组, 并将各位设置为指定的值
	fmt.Printf("bitfield set %v\n", rdb.BitField(ctx, bitfieldkey, "set", "i64", "#0", "99"))
	fmt.Printf("bitfield set %v\n", rdb.BitField(ctx, bitfieldkey, "set", "i64", "#1", "98"))
	fmt.Printf("bitfield set %v\n", rdb.BitField(ctx, bitfieldkey, "set", "i64", "#2", "97"))

	//todo://指定位的值的增量
	rdb.BitField(ctx, bitfieldkey, "incrby", "i64", "#0", "-2")

	fmt.Printf("获取 bitfieldkey %v\n", rdb.Get(ctx, bitfieldkey))
}

func TestDatabase() {

	fmt.Printf("获取指定的key是否存在:%v\n", rdb.Exists(ctx, bitfieldkey))
	fmt.Printf("获取指定的key类型:%v\n", rdb.Type(ctx, "list1"))

	//todo 重命名指定的key
	fmt.Printf("重命名key, 并把值迁移过去 %v\n", rdb.Rename(ctx, bitfieldkey, bitfieldkey + "_new"))
	fmt.Printf("获取 bitfieldkey %v\n", rdb.Get(ctx, bitfieldkey))
	fmt.Printf("获取 bitfieldkey %v\n", rdb.Get(ctx, bitfieldkey + "_new"))

	rdb.Set(ctx, bitfieldkey, "----", 0)
	fmt.Printf("获取选中的db的指定key值:%v\n", rdb.Get(ctx, bitfieldkey))

	//todo:移动key到指定数据库, 仅当对方数据无此key时有效
	fmt.Printf("移动指定的key到另外的db %v\n", rdb.Move(ctx, bitfieldkey, 2))

	//todo 通过ctx获取一个新的连接
	var newConn *redis.Conn = rdb.Conn(ctx)
	//todo 选择指定的数据库
	newConn.Select(ctx, 2)

	//todo 删除指定的key
	//newConn.Del(ctx, bitfieldkey)
	//rdb.Del(ctx, bitfieldkey)
	fmt.Printf("获取选中的db的指定key值:%v\n", newConn.Get(ctx, bitfieldkey))

	fmt.Printf("从当前数据库随机返回一个key:%v\n", newConn.RandomKey(ctx))

	fmt.Printf("从当前数据库返回key的数量:%v\n", newConn.DBSize(ctx))

	//todo 从当前数据库返回所有的key值, 带入正则
	fmt.Printf("从当前数据库返回所有key:%v\n", newConn.Keys(ctx, "*"))

	//todo 增量迭代返回所有的key值
	fmt.Printf("从当前数据库返回所有key:%v\n", newConn.Scan(ctx, 0, "*", 2))

	var sort *redis.Sort = &redis.Sort{
		//By           : "Score", //按照其他键的值来排序
		Offset 		 : 3, //跳过指定元素的数量, 比如要跳过前3条数据
		Count 		 :10,  //跳过之后的数据值
		//Get          : []string{""}, //根据排序好的规则, 提取出指定key的值, 比如根据uid排序, 得到id, 然后带入"u_nmae_*", 取得对应的key值 (可获取多个外部key的值)
		//Order         : "DESC", //排序规则  默认升序, 显示写入"DESC"降序
		Alpha         : true,   //对字符串进行排序
	}

	fmt.Printf("%v\n", rdb.Sort(ctx, "zset1", sort))
	fmt.Printf("获取所有的值:%v\n", rdb.ZRangeWithScores(ctx, "zset1", 0, -1))

	//todo 清空当前数据库所有数据
	//newConn.FlushDB(ctx)
	fmt.Printf("清空后, 获取所有数据:%v\n", newConn.Keys(ctx, "*"))
	//newConn.FlushAll(ctx)

	//todo 交换两个数据库的数据
	fmt.Printf("交换两个数据库:%v\n", newConn.SwapDB(ctx, 2, 1))
	newConn.Select(ctx, 1)
	fmt.Printf("清空后, 获取所有数据:%v\n", newConn.Keys(ctx, "*"))
	newConn.Select(ctx, 2)
	fmt.Printf("清空后, 获取所有数据:%v\n", newConn.Keys(ctx, "*"))


	//todo 指定过期时间
	//rdb.Expire(ctx, "key1", time.Second * 2)
	//rdb.ExpireAt(ctx, "key1", time.Now())

	//todo 获取过期时间
	//rdb.TTL(ctx, "key1")
	//rdb.PExpire()
}

func TestTransaction() {

}

func TestRedisLua(){
	//rdb.Eval()
	//rdb.ScriptFlush()
	//rdb.EvalSha()
	//rdb.ScriptLoad()
	//rdb.ScriptExists()
	//rdb.ScriptKill()


	//todo redis持久化  手动保存数据
	rdb.Save(ctx)
	rdb.BgSave(ctx)

	//后台保存, 创建一个aod的优化版本. 成功与否不会影响原来的数据
	rdb.BgRewriteAOF(ctx)
	rdb.LastSave(ctx)
}

func TestPublishAndSubcribe() {
	//todo 订阅频道
	sub := rdb.Subscribe(ctx, "ss")
	//todo 订阅给定条件的多个频道
	//pubsub := rdb.PSubscribe(ctx, "ss", "22")

	//todo 取消订阅
	//sub.Unsubscribe(ctx, "ss")

	iface, _ := sub.Receive(ctx)
	ch := sub.Channel()

	//获取订阅状态
	switch iface.(type) {
	case *redis.Message:
		fmt.Printf("Message:%v\n", iface)
	case *redis.Subscription:
		// Can be "subscribe", "unsubscribe", "psubscribe" or "punsubscribe".
		fmt.Printf("Subscription:%v\n", iface)
	case *redis.Pong:
		fmt.Printf("Pong:%v\n", iface)
	default:
		fmt.Printf("error:%v\n", iface)
	}

	go func() {
		for {
			for v := range ch {
				fmt.Printf("输出发布内容:%v\n", v)
			}
		}
	}()

	//todo 发布
	for i := 0; i < 10; i++ {
		rdb.Publish(ctx, "ss", "publish + " + strconv.Itoa(i))
	}

	//todo 获取所有订阅的频道
	fmt.Printf("%v\n", rdb.PubSubChannels(ctx, "*"))

}
/*
设置或取消 当前服务器为指定服务器的从属服务器.
*/
func TestSlaveOf()  {
	//todo 将当前服务器作为指定服务器的从属服务器
	fmt.Printf("指定当前服务器为指定服务器的从属服务器: %v\n", rdb.SlaveOf(ctx, "127.0.0.1", "6380"))
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6380",
	})

	//todo 在新服务器创建一个key, 原来的rdb会自行从新服务器复制一份
	client.Del(ctx, "woxwoxwoxwoxwoxwoxwoxwox")
	rdb.Del(ctx, "woxwoxwoxwoxwoxwoxwoxwox")
	fmt.Printf("new demo1-mysql all keys: %v\n", client.Keys(ctx, "*"))
	fmt.Printf("指定服务器添加一个Key %v\n", client.Set(ctx, "woxwoxwoxwoxwoxwoxwoxwox", "1", 0))
	fmt.Printf("new demo1-mysql all keys: %v\n", client.Keys(ctx, "*"))
	fmt.Printf("original demo1-mysql all keys: %v\n", rdb.Keys(ctx, "*"))

	//todo 将当前服务器取消作为从属服务器
	fmt.Printf("指定当前服务器为指定服务器的从属服务器: %v\n", rdb.SlaveOf(ctx, "no", "one"))

	//todo 取消作为从属服务器后, 对新服务器设置新key, 原来的rad不会复制
	rdb.Del(ctx, "~~~woxwoxwoxwoxwoxwoxwoxwox")
	fmt.Printf("~~~new demo1-mysql all keys: %v\n", client.Keys(ctx, "*"))
	fmt.Printf("~~~指定服务器添加一个Key %v\n", client.Set(ctx, "~~~woxwoxwoxwoxwoxwoxwoxwox", "1", 0))
	fmt.Printf("~~~new demo1-mysql all keys: %v\n", client.Keys(ctx, "*"))
	fmt.Printf("~~~original demo1-mysql all keys: %v\n", rdb.Keys(ctx, "*"))
}

func TestRedisSystemCmd() {

	//todo 已有授权密码的情况下, 必须先使用auth授权后, 才能设置key 重设密码等操作.
	//todo 没有授权密码的情况下, 可以先设置密码, 然后使用授权连接.
	//todo rdb.Conn 内部是从baseClient的子类, 从Client参数创建
	client := rdb.Conn(ctx)
	client.Auth(ctx, "123")
	fmt.Printf("设置密码后, 验证授权是否成功%v\n", client.Ping(ctx))

	//todo 动态修改redis配置
	fmt.Printf("设置redis授权密码:%v\n", client.ConfigSet(ctx, "requirepass", "123"))
	fmt.Printf("获取redis授权密码:%v\n", client.ConfigGet(ctx, "requirepass"))
	fmt.Printf("设置授权密码后, 赋值key %v\n", client.Set(ctx, "设置密码后, 直接设置key", "11", 0))

	client2 := rdb.Conn(ctx)
	client2.Auth(ctx, "123")
	fmt.Printf("设置授权密码后, 赋值key %v\n", client2.Set(ctx, "设置密码后, 直接设置key2", "11", 0))



	//todo 获取redis服务器的各种信息
	fmt.Printf("获取redis的各种信息:%v\n", client.Info(ctx))

	//fmt.Printf("~~~original demo1-mysql all keys: %v\n", demo1-mysql.Keys(ctx, "*"))
	//keys, _ := demo1-mysql.Keys(ctx, "*").Result()
	//for _, v := range keys {
	//	demo1-mysql.Del(ctx, v)
	//}

	//todo redis内部命令 quit 无法实现, 代码中不能调用
	//demo1-mysql.Quit(ctx)
	//rdb.Quit(ctx)

	//todo 内部命令 会关闭redis-server 这里的shutdown 执行save or nosave
	//demo1-mysql.Shutdown(ctx)
	//demo1-mysql.ShutdownSave(ctx)
	//demo1-mysql.ShutdownNoSave(ctx)

	//todo 获取当前系统时间
	fmt.Printf("%v\n", client.Time(ctx))

	//todo 设置客户端名称, 不可有中文
	client.ClientSetName(ctx, "joo")
	fmt.Printf("获取客户端名称:%v\n", client.ClientGetName(ctx))

	fmt.Printf("获取连接的客户端:%v\n", client.ClientList(ctx))


	str, _ := client2.ClientList(ctx).Result()
	str = strings.TrimRight(str, "\n")
	fmt.Printf("demo1-mysql list:%q\n", str)
	slice := strings.Split(str, "\n")
	resultSlice := make([]map[string]interface{}, len(slice))
	for i, v := range slice{
		subSlice := strings.Split(v, " ")
		argMap := make(map[string]interface{}, len(subSlice))

		for _, v1 := range subSlice {
			subSlice1 := strings.Split(v1, "=")
			if len(subSlice1) == 2 {
				argMap[subSlice1[0]] = subSlice1[1]
			}
		}
		resultSlice[i] = argMap
	}

	marshalBytes, _ := json.Marshal(resultSlice)
	fmt.Printf("输出组合好的数据:%v\n", string(marshalBytes))

	ipPort := resultSlice[0]["addr"]
	//todo 杀死一个客户端  此处需要注意的是, 调用时, 如果杀死的是自己, 那么无法再次通过自己获取client list
	//todo 比如这里, client杀死的是client list中第一个就是自己, 那么下面使用client再次获取client list, 就无法获取到新数据
	if ipPort != nil {
		fmt.Printf("杀死某一个客户端:%v\n", client.ClientKill(ctx, ipPort.(string)))
	}

	//todo 这里只能使用其他client来获取连接中的客户端
	str, _ = client2.ClientList(ctx).Result()
	str = strings.TrimRight(str, "\n")
	fmt.Printf("demo1-mysql list:%q\n", str)

	//todo 客户端被关闭后, 无法继续操作
	client.Set(ctx, "qwertttt", 11, 0)
	fmt.Printf("%v\n", client.Get(ctx, "qwertttt"))


	//todo 重置redis系统的配置
	//client2.ConfigResetStat(ctx)
	//todo 将config set修改的内容, 写入到redis.conf中去. 如果启动服务器时, 没有载入redis.conf, 返回错误  该操作是原子性的
	fmt.Printf("重写redis.conf %v\n", client2.ConfigRewrite(ctx))
	fmt.Printf("获取redis所有配置: %v\n", client2.ConfigGet(ctx, "*"))



	//todo 指定一个带redis.conf的server启动后, 修改其密码, 并复写redis.conf
	client3 := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6381",
	}).Conn(ctx)

	client3.Auth(ctx, "")
	client3.ConfigSet(ctx, "requirepass", "")
	fmt.Printf("重写redis.conf %v\n", client3.ConfigRewrite(ctx))
}

func TestBuildinDebug() {
	client := rdb.Conn(ctx)
	client.Auth(ctx, "123")
	fmt.Printf("ping %v\n", client.Ping(ctx))

	//todo 调试打印 echo
	fmt.Printf("echo %v\n", client.Echo(ctx, "11"))

	//todo 观察redis内部key的对象
	client.LPush(ctx, "list", "11, 22, 33", "abc")
	fmt.Printf("%v\n", client.Keys(ctx, "*"))
	fmt.Printf("debug object:%v\n", client.DebugObject(ctx, "list"))
	fmt.Printf("object encoding:%v\n", client.ObjectEncoding(ctx, "list"))

	//todo 配置对查询超过100微秒的数据进行记录, 以及设置最多存储多少条
	client.ConfigSet(ctx, "slowlog-log-slower-than", "100")
	client.ConfigSet(ctx, "slowlog-max-len", "200")

	//todo 获取slowlog
	client.SlowLogGet(ctx, 11)

	//todo redis内部命令 "MONITOR", 实时打印redis执行的命令

	//todo 迁移一个字段到另一个server, 迁移后, 本server key移除. 如果对方server包含此key或有密码, 都无法迁移成功
	client2 := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6381",
	}).Conn(ctx)
	//client2.Del(ctx, "migrate_set_key")
	//demo1-mysql.Del(ctx, "migrate_set_key")

	client.SAdd(ctx, "migrate_set_key", "aaaaa")
	fmt.Printf("迁移字段到另一个实例:%v\n", client.Migrate(ctx, "127.0.0.1", "6381", "migrate_set_key", 0, 0))
	fmt.Printf("获取迁移前client是否包含字段:%v\n", client.SIsMember(ctx, "migrate_set_key", "被迁移的字段"))
	fmt.Printf("获取迁移后client2是否包含字段:%v\n", client2.SIsMember(ctx, "migrate_set_key", "被迁移的字段"))

	//todo 序列化指定的key的值
	dumpCmd := client.Dump(ctx, "migrate_set_key")
	str, _ := dumpCmd.Result()
	fmt.Printf("Dump: %v\n", str)
	//demo1-mysql.Del(ctx, "migrate_set_key")

	//todo 恢复序列化后的值到指定的key, 如key已存在, 失败
	restoreCmd := client.Restore(ctx, "migrate_set_key1", 0, str)
	fmt.Printf("Restore: %v\n", restoreCmd)
	//todo 恢复序列化后的值到指定的key, 强制替换
	restoreCmd = client.RestoreReplace(ctx, "migrate_set_key", 0, "2")
	fmt.Printf("RestoreReplace: %v\n", restoreCmd)

	fmt.Printf("获取指定的set的并集:%v\n", client.SUnion(ctx, "migrate_set_key", "migrate_set_key"))

	//todo SYNC 全量同步, PSYNC 部分同步
	//todo 全量同步：slave启动或者slave断开重连master的时候，slave会发生SYNC命令给master，master接收到该命令后，
	//todo 则会通过bgsave启动一个子进程将当前时间点的master全部数据的快照写到一个文件中，
	//todo 然后发送给slave。slave接收到之后则清空内存，载入这些数据。

	//todo  内部没有实现, 会释放一个panic. PSYNC 该库没有实现该命令
	//demo1-mysql.Sync(ctx)

	//todo 手动保存redis数据
	// rdb 需要在redis.conf 配置, 比如 save 60 100  (60s内100个键被改动时保存)
	// aof 需要在redis.conf 配置appendonly yes
	// 没有指定conf的server保存路径为用户目录下面, 指定conf的server保存在指定路径下面(macos 默认指定路径/usr/local/var/db/redis)
	fmt.Printf("开启AOF:%v\n", client2.ConfigSet(ctx, "appendonly", "yes"))
	fmt.Printf("重写配置到conf文件:%v\n", client2.ConfigRewrite(ctx))
	fmt.Printf("手动保存 RDB (Redis 快照)%v\n", client2.Save(ctx))
	fmt.Printf("手动保存 AOF (Redis 执行代码%v\n)", client2.BgRewriteAOF(ctx))
}
