package codec

import "io"

//通用头
type Header struct {
	seq           uint64
	servicemethod string
	codec         *Codec
}

// 接口化 可以被多种编码解码方式实现
// 定义方法 ： 编码、解码方法
// 方法参数 ： 请求头、请求体，响应体
type Codec interface {
	io.Closer
	//encode(writer io.Writer)
	//decode(reader io.Reader)
	// 为 header body 定义方法
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	WriteHeader(*Header) error
	WriteBody(interface{}) error
}

// Codec 接口被实现后，就可以直接调用接口来使用方法了。
