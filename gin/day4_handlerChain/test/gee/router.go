package gee

import (
	"net/http"
	"strings"
)

type Handler func(ctx *Context)

type Router struct {
	nodeMap         map[string]*Node
	mapPathHandlers map[string]Handler
}

//包级别引用
func newRouter() *Router {
	return &Router{
		nodeMap:         make(map[string]*Node),
		mapPathHandlers: make(map[string]Handler),
	}
}

func parsePath(path string) []string {
	vs := strings.Split(path, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

//路由映射
func (r *Router) addRouter(method, url string, handler Handler) {
	parts := parsePath(url)
	key := method + "-" + url
	if r.nodeMap[method] == nil {
		r.nodeMap[method] = &Node{}
	}
	//插入节点： 从跟节点开始 height=0
	r.nodeMap[method].buildTree(parts, url, 0)
	r.mapPathHandlers[key] = handler
}

func (r *Router) getRoute(method, path string) (*Node, map[string]string) {
	searchParts := parsePath(path)
	params := make(map[string]string)
	root, ok := r.nodeMap[method]
	if !ok {
		return nil, nil
	}
	n := root.checkTree(searchParts, 0)
	if n != nil {
		parts := parsePath(n.path)
		for index, part := range parts {
			//动态路由的：后步，如：abc ,作为params的key,
			//value是路径。
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			//如果是通配符
			//通配符后内容作为key,value是？
			if part[0] == '*' {
				params[part[1:]] = strings.Join(
					searchParts[index:], "/")
				break
			}
		}
	}
	return n, params
}

//将从路由匹配得到的Handler添加到c.handlers列表中，执行c.next()
func (r *Router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.path
		c.handlers = append(c.handlers, r.mapPathHandlers[key])
	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {
			c.RespByString(http.StatusNotFound, "404 NOT FOUND :%s\n", c.Path)
		})
	}
	c.Next()
}
