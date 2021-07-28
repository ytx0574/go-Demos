package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

/*
客户端实现远程调用  比如下面得
*/

func main() {

	loc, _ := time.LoadLocation("")
	time.Now().In(loc).Date()

	log.Printf("我去你大爷的\n")
	fmt.Printf("哈哈哈哈\n")

	cli , err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	args := &Args{A: 10, B:20}
	var reply int
	err = cli.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("arith: %v * %v = %v\n", args.A, args.B, reply)

	log.Ldate
}