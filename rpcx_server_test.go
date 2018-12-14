package xserver

import (
	"fmt"
	"testing"
)

func TestRpcxServer(t *testing.T) {
	grpcServer, err := NewServerRPCX(Config{Address: []string{"127.0.0.1:8973"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer.RegisterName("Logger", "", "")
	err = grpcServer.Listen()
	defer grpcServer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
