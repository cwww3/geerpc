package main

import (
	"geerpc/common"
	"geerpc/server"
	"log"
	"net"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func main() {
	var foo Foo
	if err := server.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}
	//	// pick a free port
	l, err := net.Listen("tcp", common.Addr)
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	server.Accept(l)
}

