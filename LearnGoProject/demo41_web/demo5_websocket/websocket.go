package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

func Echo(conn *websocket.Conn) {

	var err error
	for {
		var replay string
		if err = websocket.Message.Receive(conn, &replay); err != nil {
			fmt.Printf("接受消息失败, %v\n", err)
			break
		}

		fmt.Printf("读取到websocket消息, %v\n", replay)

		msg := "received: " + replay

		if err = websocket.Message.Send(conn, msg); err != nil {
			fmt.Printf("发送消息失败 %v\n", err)
			break
		}
	}
}

func main() {
	//http.Handle("/", websocket.Handler(Echo))
	//
	//http.ListenAndServe(":1234", nil)


	r, err := http.NewRequest("GET", "http://47.56.9.229:8008", nil)
	if err != nil {
		log.Printf("初始化请求失败:%v\n", err)
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("请求失败 %v\n", err)
	}else {
		defer  res.Body.Close()
		info, err := ioutil.ReadAll(res.Body)
		log.Printf("请求成功:%v\nerr:%v\n", string(info), err)
	}
}
