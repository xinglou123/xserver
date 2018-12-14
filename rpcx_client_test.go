package xserver

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	param := map[string]interface{}{}
	param["path"] = ""
	param["page"] = 1
	param["limit"] = 10

	req := Request{
		ServicePath:   "Logger",
		ServiceMethod: "List",
		Params:        param,
	}

	for a := 0; a < 20; a++ {
		start := time.Now()
		rpcx, _ := NewClientRPCX(Config{Address: []string{"127.0.0.1:8972"}, ServicePath: "Logger"})
		defer rpcx.Close()
		reply, _ := rpcx.Call(req)
		latency := fmt.Sprintf("%dms", time.Since(start).Nanoseconds()/1e6)
		fmt.Println(latency)
		fmt.Println(reply)
	}

}
