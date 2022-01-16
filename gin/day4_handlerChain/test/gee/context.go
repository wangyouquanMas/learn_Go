package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	//req info
	Method string
	Path   string
	//resp info
	StatusCode int
	Params     map[string]string
	//middleware info
	handlers []Handler
	index    int
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.RespByJSON(code, H{"message": err})
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Method: r.Method,
		Path:   r.URL.Path,
		index:  -1,
	}
}

//TODO  Next设计
//这里 index是 context指针的值，因此类似于全局变量
//例1 ： 所以假设有3个中间件A，B，C，一个业务handler D。
//假设执行顺序是  A,B，C D，那么执行完D, index已经变成了4，退出for循环，所以C 的next执行结束。
//继续执行C的后续逻辑。 执行结束，表明B.next执行结束，继续执行B的后续逻辑...

//例2 : 假设 上述ABCD中，C不需要执行 next,只作用于请求前。
//那么 A->B->C  C执行结束，此时index为3，而s为4，会继续执行for循环， 此时发现c.handlers[3]就是业务handler。 D执行结束，index为5 ，此时退出next，继续执行B的next()后逻辑。。。

//那为什么不能这样设计？
//如例2 ： 当 C执行结束后， index为2，退出next，继续执行B后续逻辑，导致 D无法自动执行。
func (c *Context) Next1() {
	c.index++
	c.handlers[c.index](c)
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Param(key string) string {
	return c.Params[key]
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
