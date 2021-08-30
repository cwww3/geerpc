package client

import (
	"fmt"
	"geerpc/server"
	"net"
	"os"
	"runtime"
	"testing"
)

func TestXDial(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		ch := make(chan struct{})
		addr := "/tmp/geerpc.sock"
		go func() {
			_ = os.Remove(addr)
			l, err := net.Listen("unix", addr)
			if err != nil {
				t.Fatal("failed to listen unix socket")
			}
			ch <- struct{}{}
			server.Accept(l)
		}()
		<-ch
		_, err := XDial("unix@" + addr)
		_assert(err == nil, "failed to connect unix socket")
	}
}