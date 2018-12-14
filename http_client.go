package xserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type ClientHttp struct {
	client *http.Client
	conf   Config
}

func NewClientHttp(cfg Config) (c ClientHttp, err error) {
	if len(cfg.Address) == 0 {
		err = errors.New("Address must be non-empty")
	}

	c.client = &http.Client{}

	c.conf = cfg

	return
}
func (c *ClientHttp) Call(ctx context.Context, req Request) (reply *Response, err error) {
	fmt.Println(req)
	if c.client != nil {
		req, err := http.NewRequest("POST", c.conf.Address[0]+"/"+c.conf.ServicePath, nil)
		if err != nil {
			return nil, errors.New("failed to create POST request")
		}
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, errors.New("request failed")
		}
		reply = new(Response)
		reply.Code = resp.StatusCode
		reply.Data = resp
	}
	return
}

func (c *ClientHttp) Close() {
	return
}
