// geektutu.com
// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*//gin框架初始化的流程
1.初始化engine
顺便完成了
	1.1注册中间件
	1.2注册路由

//响应流程
1.路由，找到handle
2.将请求和响应用Context包装起来供业务代码使用   [!!!!]
3.依次调用中间件和处理函数
4.输出结果*/

func main() {
	//gin.Default()生成了一个实例，这个实例即 WSGI [Web Server Gateway Interface]应用程序。
	//.Engine gin 引擎，是框架的实例，它包含多路复用器，中间件和配置设置
	r := gin.Default()
	//使用r.Get("/", ...)声明了一个路由 【有多种、get\post\put等等】，告诉 Gin 什么样的URL 能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息。
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	//	数解析 : 动态路由 [当每个url携带了不同的请求参数]
	r.GET("/user/:name",func(c *gin.Context){
		name:= c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})


	// 获取query参数 如 curl "http://localhost:9999/users?name=Tom&role=student"
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role","teacher")
		c.String(http.StatusOK,"%s is a %s",name,role)
	})

	// 获取post参数
	r.POST("/form",func(c *gin.Context){
		username := c.PostForm("username")
		password := c.DefaultPostForm("password","0000000")

		c.JSON(http.StatusOK,gin.H{
			"username": username,
			"password": password,
		})
	})


	// GET 和 POST 混合
	//$ curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	//{"id":"9876","page":"7","password":"1234","username":"geektutu"}
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})


	//Map参数(字典参数)
	//curl -g "http://localhost:9999/post?ids[jack]=001&ids[tom]=002" -d "names[lucy]=a&names[burry]=b"
	r.POST("/post",func(c *gin.Context){
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK,gin.H{
			"ids":ids,
			"names":names,
		})
	})


// 后续参考 https://geektutu.com/post/quick-go-gin.html

	//r.Run()函数来让应用运行在本地服务器上，默认监听端口是 _8080_，可以传入参数设置端口，例如r.Run(":9999")即运行在 _9999_端口。
	r.Run(":9999") // listen and serve on 0.0.0.0:8080
}
