package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//解析request

//URL *url.URL : 请求行中url

//如何获取req请求体中的数据
func GetHttpBody(w http.ResponseWriter, r *http.Request) {
	// ContentLength :请求体中字节大小
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

//获取Form : 包含url中query数据和 表单数据
//postform:只能获取表单数据
func FormPostForm(w http.ResponseWriter, r *http.Request) {

	//ParseForm : 只有先解析后，才能获取Form数据
	r.ParseForm()

	fmt.Fprintln(w, "r.PostForm->:", r.PostForm)
	fmt.Fprintln(w, "r.Form->:", r.Form)
}

//获取multipleform
func MultiPartform(w http.ResponseWriter, r *http.Request) {
	//指定可存储在内容中的最大字节数，剩余的字节会被存储到磁盘中
	r.ParseMultipartForm(1024)

	fileHeader := r.MultipartForm.File["file_to_upload"][0]
	f, err := fileHeader.Open()

	if err != nil {
		fmt.Fprintln(w, "fileHeader.Open()->", err.Error())
	} else {
		//？file实现了read方法
		dataFromFile, err := ioutil.ReadAll(f)
		if err == nil {
			fmt.Fprintf(w, "%s\n", string(dataFromFile))
		}
	}
	fmt.Fprintln(w, "r.MultipartForm->", r.MultipartForm)
}

func main() {
	//h1 := func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	//	fmt.Fprintln(w, r.URL.RawQuery)
	//	fmt.Fprintln(w, r.URL.Host)
	//	fmt.Fprintln(w, r.URL.Path)
	//	fmt.Fprintln(w, "通过handleFunc启动一个http服务")
	//	rawQuery := r.URL.RawQuery
	//	//ParseQuery :解析请求返回 map 类型 的key-value对
	//	va, _ := url.ParseQuery(rawQuery)
	//	fmt.Fprintln(w, "获取url中key为name值", va.Get("name"))
	//}
	//h2 := func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "Hello from a HandleFunc #2!\n")
	//	fmt.Fprintln(w, r.URL.RawQuery)
	//	fmt.Fprintln(w, r.URL.Host)
	//	fmt.Fprintln(w, r.URL.Path)
	//
	//	for key := range r.Header {
	//		fmt.Fprintf(w, "%s:%s\n", key, r.Header[key])
	//	}
	//
	//	fmt.Fprintln(w, "获取Accept-Language的val值", r.Header["Accept-Language"])
	//
	//	fmt.Fprintln(w, "通过handleFunc启动一个http服务")
	//}

	//http.HandleFunc("/", h1)
	//http.HandleFunc("/endpoint", h2)
	//http.HandleFunc("/getBody", GetHttpBody)
	//http.HandleFunc("/formpost", FormPostForm)
	http.HandleFunc("/multiPartform", MultiPartform)
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// handler设置为nil，默认为多路处理器作为handler，默认多路处理器实现了ServerHttp方法
	log.Fatal(http.ListenAndServe(":8080", nil))
}
