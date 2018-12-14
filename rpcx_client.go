package xserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/client"
)

type ClientRPCX struct {
	mclient client.XClient
}

func NewClientRPCX(cfg Config) (c ClientRPCX, err error) {

	if len(cfg.Address) == 0 {
		err = errors.New("address must be specified")
		return
	}
	var d client.ServiceDiscovery
	if len(cfg.Address) > 1 {
		var pairs []*client.KVPair
		for _, v := range cfg.Address {
			pairs = append(pairs, &client.KVPair{Key: v})
		}
		d = client.NewMultipleServersDiscovery(pairs)
	} else {
		d = client.NewPeer2PeerDiscovery("tcp@"+cfg.Address[0], cfg.Meta)
	}
	if len(cfg.ServicePath) == 0 {
		err = errors.New("ServicePath must be specified")
		return
	}
	c.mclient = client.NewXClient(cfg.ServicePath, client.Failtry, client.RandomSelect, d, client.DefaultOption)

	return
}

func (c *ClientRPCX) CallWithContext(ctx context.Context, req Request) (reply *Response, err error) {
	if c.mclient != nil {
		reply = new(Response)
		err = c.mclient.Call(ctx, req.ServiceMethod, req.Params, reply)
	} else {
		err = errors.New("client.XClient is nil")
	}
	return
}

func (c *ClientRPCX) Call(req Request) (reply *Response, err error) {
	fmt.Println(req)
	if c.mclient != nil {
		reply = new(Response)
		err = c.mclient.Call(context.Background(), req.ServiceMethod, req.Params, reply)
	} else {
		err = errors.New("client.XClient is nil")
	}
	return
}

func (c *ClientRPCX) Close() {
	if c.mclient != nil {
		c.mclient.Close()
	}
}
