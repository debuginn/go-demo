package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立 utp 服务器
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		fmt.Printf("listen failed error:%v\n", err)
		return
	}
	defer listen.Close() // 使用完关闭服务

	for {
		// 接收数据
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read data error:%v\n", err)
			return
		}
		fmt.Printf("addr:%v\t count:%v\t data:%v\n", addr, n, string(data[:n]))
		// 发送数据
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("send data error:%v\n", err)
			return
		}
	}

}
