package xserver

import (
	"errors"
	"net/http"
	"strconv"
)

type ServerH struct {
	server      *http.Server
	certificate string
	key         string
}

type ServerHConfig struct {
	Port        int
	Certificate string
	Key         string
}

func NewServerH(cfg ServerHConfig) (s ServerH, err error) {
	if cfg.Port == 0 {
		err = errors.New("Port must be non-zero")
		return
	}

	if cfg.Certificate == "" {
		err = errors.New("Certificate must be specified")
		return
	}

	if cfg.Key == "" {
		err = errors.New("Key must be specified")
		return
	}

	s.server = &http.Server{
		Addr: ":" + strconv.Itoa(cfg.Port),
	}

	s.certificate = cfg.Certificate
	s.key = cfg.Key

	http.HandleFunc("/upload", s.Upload)

	return
}

func (s *ServerH) Listen() (err error) {
	err = s.server.ListenAndServeTLS(
		s.certificate, s.key)
	if err != nil {
		err = errors.New("failed during server listen and serve")
		return
	}

	return
}

func (s *ServerH) Upload(w http.ResponseWriter, r *http.Request) {
	//var (
	//	err           error
	//	bytesReceived int64 = 0
	//	buf                 = new(bytes.Buffer)
	//)
	//
	//bytesReceived, err = io.Copy(buf, r.Body)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprint(w, "%+v", err)
	//	return
	//}

	return
}

func (s *ServerH) Close() {
	return
}
