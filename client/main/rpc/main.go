package main

import (
	"context"
	"geerpc/client"
	"geerpc/common"
	"log"
	"sync"
	"time"
)

type Args struct{ Num1, Num2 int }

func main() {
	c, _ := client.Dial("tcp", common.Addr)
	defer func() { _ = c.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			ctx, _ := context.WithTimeout(context.Background(), time.Second)
			if err := c.Call(ctx, "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait()
}
