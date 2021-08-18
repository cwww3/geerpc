package main

import (
	"geerpc"
	"geerpc/common"
	"log"
	"net"
)

func main() {
	//	// pick a free port
		l, err := net.Listen("tcp", common.Addr)
		if err != nil {
			log.Fatal("network error:", err)
		}
		log.Println("start rpc server on", l.Addr())
		geerpc.Accept(l)
}
