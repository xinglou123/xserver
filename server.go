package xserver

type Server interface {
	Register(meta string, obj interface{}) (err error)
	RegisterName(name, meta string, obj interface{}) (err error)
	Listen() (err error)
	Close()
}
