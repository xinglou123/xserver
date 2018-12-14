package xserver

type Response struct {
	Code int         `json:"code"` // return code, 0 for succ
	Msg  string      `json:"msg"`  // message
	Data interface{} `json:"data"` // data object
}
