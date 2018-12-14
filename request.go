package xserver

type Request struct {
	/***
	服务名
	*/
	ServicePath string
	/***
	服务方法
	*/
	ServiceMethod string
	/***
	接口参数
	*/
	Params map[string]interface{}
}
