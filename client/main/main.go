package main

import (
	"fmt"
	"geerpc/client"
	"geerpc/common"
	"log"
	"sync"
)

func main() {
	c, _ := client.Dial("tcp", common.Addr)
	defer func() { _ = c.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := c.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
