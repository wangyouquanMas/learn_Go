package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	//req info
	Method string
	Path   string
	//resp info
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}

//请求
//获取请求信息
//1获取请求参数
func (c *Context) GetQueryParmatersByForm(key string) string {
	return c.Req.FormValue(key)
}

//2获取请求路由
func (c *Context) GetUrlInfo(key string) string {
	return c.Req.URL.Query().Get(key)
}

//响应
func (c *Context) WriterStatusCode(code int) {
	c.Writer.WriteHeader(code)
}

//封装：封装不应该直接调用最底层的方法，应该将它封装起来
//比如让结构体，context去调用 header set方法就很麻烦，要写一大串内容
func (c *Context) WriterHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

//多种返回类型
func (c *Context) RespByJSON(code int, obj interface{}) {
	c.WriterHeader("Content-Type", "application/json")
	//c.Writer.Header().Set("Content-Type", "application/json")
	c.WriterStatusCode(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) RespByString(code int, format string, value ...interface{}) {
	c.WriterHeader("Content-Type", "text/plain")
	//c.Writer.Header().Set("Content-Type", "application/json")
	c.WriterStatusCode(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) RespByHTML(code int, html string) {
	c.WriterHeader("Content-Type", "text/html")
	//c.Writer.Header().Set("Content-Type", "application/json")
	c.WriterStatusCode(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) RespByByte(code int, data []byte) {
	c.WriterStatusCode(code)
	c.Writer.Write(data)
}
