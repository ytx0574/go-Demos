
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())


	data := make([]byte, 1024)

	i := ""
	for {
		 count, addr, err := conn.ReadFromUDP(data)
		 if err == nil {
		 	log.Printf("获得服务端反馈:%v %v %v\n", count, addr, string(data[:count]))
		 }
		 i += " 1"
		 conn.Write([]byte(i))

		if len(i) > 10 {
			break
		}
	}
}