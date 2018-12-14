package xserver

import (
	"context"
)

type Client interface {
	CallWithContext(ctx context.Context, req Request) (reply *Response, err error)
	Call(req Request) (reply *Response, err error)
	Close()
}
