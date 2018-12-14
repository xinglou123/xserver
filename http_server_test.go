package xserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestHttpServer(t *testing.T) {
	httpserver, err := NewServerHttp(Config{Address: []string{":20000"}, ReadTimeout: 30, WriteTimeout: 30})
	if err != nil {
		fmt.Println(err)
		return
	}
	agin := gin.New()
	httpserver.Router = agin
	err = httpserver.Listen()
	defer httpserver.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
