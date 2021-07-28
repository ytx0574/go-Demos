package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/clientv3"
	pb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	"sort"
	"strings"
	"time"
	"unsafe"
)

var ctx = context.Background()

func main() {
	//sort.IntSlice{}
	//sort.Reverse()
	x := 10
	slice := []int{1, 2, 3, 4, 5, 6, 7, 10, 15}
	//todo 必须是在一个有序的序列中查找. 找不到的情况下会返回n 不是-1
	//todo 传入的n <= len(slice)
	n := sort.Search(len(slice), func(i int) bool {
		fmt.Println(i, slice[i])
		return slice[i] >= x
	})
	fmt.Println(n, "111")
	if n < len(slice) && slice[n] == x {
		fmt.Printf("查找的数值在slice中\n")
	}



	//ClientV3Test()
clientv3.
	cli, err := client.New(client.Config{
		Endpoints:               strings.Split("http://127.0.0.1:2379", ","),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: 10 * time.Second,
	})
	if err != nil {
		fmt.Errorf("demo1-mysql.New err: %v", err)
	}

	fmt.Printf("cli = %v\n", cli)


	keyapi := client.NewKeysAPI(cli)
	//_, err = keyapi.Set(ctx, "/service.micro-mall-users.default", "222", nil)
	//if err != nil {
	//	fmt.Printf("set err:%v\n", err)
	//	return
	//}

	getRes, err := keyapi.Get(ctx, "11", nil)
	if err != nil {
		fmt.Printf("get err:%v\n", err)
		return
	}

	fmt.Printf("res = %v\n", getRes)
	fmt.Println(getRes.Node.Key, getRes.Node.Value)



}

func ClientV3Test()  {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()

	if err == nil {
		fmt.Printf("输出etcd demo1-mysql:%bv\n", cli)
	}else {
		return
	}

	kv := clientv3.NewKV(cli)
	putRes, err := kv.Put(ctx, "/test/01", "test01 value")
	//todo 此处注意, putRes内部对象为PutResponse, 但实际指向的是pb.PutResponse
	putResO := (*pb.PutResponse)(unsafe.Pointer(putRes))
	if err == nil {
		fmt.Printf("put /test/01 result: %v header:%v, prev value:%v\n", putRes.Header, putRes.PrevKv, putRes.PrevKv)
		fmt.Printf("get prev value:%v\n", putResO.GetPrevKv())
	}

	kv.Put(ctx, "/test/02", "test02 value")
	kv.Put(ctx, "/testss", "ss")


	getRes, err := kv.Get(ctx, "/testss")
	if err == nil {
		fmt.Printf("get res value:%v\n", getRes)
	}else {
		fmt.Printf("get err:%v\n", err)
	}

	//todo Put Get都可以带入额外的OP. 增加附加条件
	getRes2, err := kv.Get(ctx, "/test", clientv3.WithPrefix())
	fmt.Printf("get res2 value:%v\n", getRes)
	fmt.Printf("get res2 first key:%v, first value:%v\n", string(getRes2.Kvs[0].Key), string(getRes2.Kvs[0].Value))


	//todo 添加一个租约, 设置xx秒过期 类似redis的TTl
	lease := clientv3.NewLease(cli)
	leaseGrantRes, err := lease.Grant(ctx, 2)

	kv.Put(ctx, "/test/03", "test03 value", clientv3.WithLease(leaseGrantRes.ID))
	//todo 提前释放一个租约, 后续无法获取值
	//lease.Revoke(ctx, leaseGrantRes.ID)
	//todo 获取所有租约, 注意是是cli还是lease
	leaseLeasesRes, err := lease.Leases(ctx)
	fmt.Printf("获取etcd中的所有租约%v\n", leaseLeasesRes)

	//todo 关闭所有的租约  如带有租约的字段关闭后再次访问会报错
	//cli.Close()

	//time.Sleep(time.Second)
	//getRes3, err := kv.Get(ctx, "/test/03")
	//fmt.Printf("getres3 value 一:%v\n", getRes3)
	//
	//////todo 对即将过期的键值续约一次, 保证下面的休眠后能获取到值
	//lease.KeepAliveOnce(ctx, leaseGrantRes.ID)
	//
	////todo Get在获取不到的数据值, err也是nil, 需自行判断内部的count 或 kvs
	//time.Sleep(time.Second * 2)
	//getRes3, err = kv.Get(ctx, "/test/03")
	//fmt.Printf("getres3 value 二:%v\n", getRes3)
	//
	////这里超过续约后的时间 无法取值
	//time.Sleep(time.Second * 2)
	//getRes3, err = kv.Get(ctx, "/test/03")
	//fmt.Printf("getres3 value 三:%v\n", getRes3)

	//todo 建立一系列op 执行一组操作
	ops := []clientv3.Op{
		clientv3.OpPut("新建一个key", "值"),
		clientv3.OpGet("新建一个key"),
		clientv3.OpDelete("新建一个key"),
		clientv3.OpGet("新建一个key"),
	}

	for _, op := range ops {
		if res, err := cli.Do(ctx, op); err == nil {
			if op.IsGet() {
				fmt.Printf("op res get:%v\n", res.Get())
			}else if op.IsPut() {
				fmt.Printf("op res put:%v\n", res.Put())
			}else if op.IsDelete() {
				fmt.Printf("op res del:%v\n", res.Del())
			}else if op.IsTxn() {
				fmt.Printf("op res txn:%v\n", res.Txn())
			}

		}
	}

	//todo txn事物. 使用Compare带入条件, 操作符result可以为"= < > !="
	//todo: func CreateRevision(key string) Cmp：key=xxx的创建版本必须满足…  比较的值是int
	//todo: func LeaseValue(key string) Cmp：key=xxx的Lease ID必须满足… 比较的值是int
	//todo: func ModRevision(key string) Cmp：key=xxx的最后修改版本必须满足… 比较的值是int
	//todo: func Value(key string) Cmp：key=xxx的创建值必须满足…    比较的值必须是 string
	//todo: func Version(key string) Cmp：key=xxx的累计更新次数必须满足… 比较的值是int

	key, val := "ssr", "val"
	keyRes, err := cli.Get(ctx, key)
	fmt.Printf("get key 一:%v\n", keyRes)

	txnRes, err := kv.Txn(ctx).If(
		clientv3.Compare(clientv3.Value(key), "=" , val),
		clientv3.Compare(clientv3.Version(key), "=", 1),
	).Then(
		clientv3.OpDelete(key),   //满足条件, 执行删除操作
	).Else(
		clientv3.OpPut(key, val), //不满足条件 执行添加操作
		clientv3.OpGet(key),
	).Commit()
	fmt.Printf("txn res:%v\n", txnRes)

	keyRes, err = cli.Get(ctx, key)
	fmt.Printf("get key 二:%v\n", keyRes)


	//todo 监听某个字段的变化
	watchCHAN := cli.Watch(ctx, key)
	go func() {
		for res := range watchCHAN{
			fmt.Printf("监听到key值发生变化:%v\n", res)
			keyVal := res.Events[0]
			fmt.Printf("监听到key的值发生变化, now:%q, before:%q\n", keyVal.Kv, keyVal.PrevKv)
		}
	}()

	kv.Put(ctx, key, "new value")
	kv.Put(ctx, key, "new value1")
	kv.Put(ctx, key, "new value2")
	kv.Put(ctx, key, "new value3")


	//todo 压缩空间, 不太懂
	compactRes, err := kv.Compact(ctx, 30)
	fmt.Printf("compact res:%v\n", compactRes)
}