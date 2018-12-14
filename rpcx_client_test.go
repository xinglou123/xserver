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
		ServicePath:   "logger",
		ServiceMethod: "List",
		Params:        param,
	}
	rpcx, _ := NewClientRPCX(Config{Address: []string{"127.0.0.1:8973"}, ServicePath: "Logger"})
	defer rpcx.Close()
	for a := 0; a < 20; a++ {
		start := time.Now()
		reply, _ := rpcx.Call(req)
		latency := time.Since(start).Seconds()
		fmt.Println(latency)
		fmt.Println(reply)
	}

}
