package gee

import (
	"net/http"
	"strings"
)

type (
	RouterGroup struct {
		prefix      string
		parent      *RouterGroup
		middlewares []Handler
		//想要直接调用addrouter， engine实现了，可以直接调用这个engine
		engine *Engine
	}
	//Engine最大的封装【就像包含router一样，同样包含routerGroup】
	Engine struct {
		router *Router
		*RouterGroup
		groups []*RouterGroup
	}
)

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = append(engine.groups, engine.RouterGroup)
	return engine
}

//定义use函数，将中间件应用到某个group
func (group *RouterGroup) Use(middlewares ...Handler) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//group实例化: 前缀，父节点，分组包含engine扩展group能力，engine有的，group都有
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	//在main函数中 先实例化了engine 通过New() ，然后routergroup中包含的是 engine指针，因此该属性也被实例化
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	//存储所有的分组开头
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

//封装原则1
func (group *RouterGroup) AddRouter(method, comp string, handler Handler) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, handler)
}

func (group *RouterGroup) Get(path string, handler Handler) {
	group.AddRouter("GET", path, handler)
}

func (group *RouterGroup) POST(path string, handler Handler) {
	group.AddRouter("POST", path, handler)
}

//engine实现了handler接口中的servehttp，所以也是handler
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

//中间件在r.use时放入 middlewares数组中，最后赋值给context.handlers
//context 包含的业务handler，在handle函数中执行，最后将从路由匹配得到的Handler添加到c.handlers列表中，执行c.next()
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []Handler
	//接收到具体请求时，判断该请求适用于哪些中间件
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := NewContext(w, r)
	//得到中间件列表后，赋值给 c.handlers
	context.handlers = middlewares
	e.router.handle(context)
}
