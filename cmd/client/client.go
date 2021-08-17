package main

import (
	"encoding/json"
	"fmt"
	"geerpc"
	"geerpc/cmd"
	"geerpc/codec"
	"log"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", cmd.Addr)
	defer func() {
		_ = conn.Close()
	}()
	time.Sleep(time.Second)
	// send options
	// 向连接中写入标记和编码方式
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	// 采用gob编码
	cc := codec.NewGobCodec(conn)
	// send request & receive response
	for i := 0; i < 2; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		// 向连接中写入标记
		_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
		// 向连接中写入head和body
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))

		//FIXME 服务端是先写入head再写入body 客户端要按顺序读

		//从连接中读取返回的head
		_ = cc.ReadHeader(h)
		log.Println("head:", h)


		// 从连接中读取返回的body
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}

