package xserver

import (
	"errors"
	"github.com/smallnest/rpcx/server"
)

type ServerRPCX struct {
	server *server.Server
	conf   Config
}

func NewServerRPCX(cfg Config) (s ServerRPCX, err error) {

	if len(cfg.Address) == 0 {
		err = errors.New("address must be specified")
		return
	}
	s.conf = cfg
	s.server = server.NewServer()
	if s.server == nil {
		err = errors.New("s.server is nil")
		return
	}
	return
}

func (s *ServerRPCX) Listen() (err error) {
	if s.server == nil {
		err = errors.New("s.server is nil")
		return
	}
	err = s.server.Serve("tcp", s.conf.Address[0])
	if err != nil {
		err = errors.New("errored listening for rpcx connections")
		return
	}
	return
}

func (s *ServerRPCX) Register(meta string, obj interface{}) (err error) {
	if s.server != nil {
		return s.server.Register(obj, meta)
	}
	return errors.New("s.server is nil")
}

func (s *ServerRPCX) RegisterName(name, meta string, obj interface{}) (err error) {
	if s.server != nil {
		return s.server.RegisterName(name, obj, meta)
	}
	return errors.New("s.server is nil")
}

func (s *ServerRPCX) Close() {
	if s.server != nil {
		s.server.Close()
	}
	return
}
