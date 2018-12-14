package xserver

type Config struct {
	Address                   []string
	Basepath                  string
	Meta                      string
	ServicePath               string
	ReadTimeout, WriteTimeout int64
}
