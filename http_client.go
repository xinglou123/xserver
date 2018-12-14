package xserver

import (
	"context"
	"errors"
	"net/http"
)

type ClientH struct {
	client  *http.Client
	address string
}

type ClientHConfig struct {
	RootCertificate string
	Address         string
}

func NewClientH(cfg ClientHConfig) (c ClientH, err error) {
	if cfg.Address == "" {
		err = errors.New("Address must be non-empty")
		return
	}

	if cfg.RootCertificate == "" {
		err = errors.New("RootCertificate must be specified")
		return
	}

	c.client = &http.Client{}

	c.address = cfg.Address

	return
}

func (c *ClientH) UploadFile(ctx context.Context, f string) (err error) {

	return
}

func (c *ClientH) Close() {
	return
}
