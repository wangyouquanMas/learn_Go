package gob

import (
	"bufio"
	"channel/gRPC/day2-client/codec"
	"encoding/gob"
	"io"
)

//1  参考gob原码 定义结构体
type GobCodec struct {
	conn io.ReadWriteCloser // ？
	buf  *bufio.Writer      // 用于提高编解码效率
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// 2 创建结构体初始化，初始化之后才能使用
func NewGobCodec(conn io.ReadWriteCloser) codec.Codec {
	buf := bufio.NewWriter(conn) // ?
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func (c *GobCodec) ReadHeader(h *codec.Header) error {
	return c.dec.Decode(h)
}

func (c *GobCodec) ReadBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *GobCodec) WriteHeader(h *codec.Header) error {
	return c.enc.Encode(h)
}

func (c *GobCodec) WriteBody(body interface{}) error {
	return c.enc.Encode(body)
}

// 父接口中嵌套子接口的方法，可以直接被父接口调用
func (c *GobCodec) Close() error {
	return c.conn.Close()
}
