package main

import (
	"fmt"
	"net"
	"network_programs/tcpnetwork/coder"
)

func ProcessConn(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := coder.Decode(conn)
		if err != nil {
			fmt.Printf("Fail to decode error [%v]\n", err)
			return
		}
		s := string(bt)
		fmt.Printf("Read from conn:[%v]\n", s)
	}
}

func main() {
	// 1. 监听端口 2.accept连接 3.开goroutine处理连接
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Fail listen.Accept:%v\n", err)
			continue
		}

		go ProcessConn(conn)
	}
}
