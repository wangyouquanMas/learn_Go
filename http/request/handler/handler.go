package handler

import (
	"channel/http/request/meta"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传html页面
		// 不能使用该路径的原因和 go.mod有关，可以使用绝对路径
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))

		//// 1 r.Method 方法用来校验请求类型
		//if r.Method == "GET" {
		//
		//	//返回登陆页
		//	//ioutil.ReadFile：读取文件，转为字节流
		//	data, err := ioutil.ReadFile("./static/view/index.html")
		//	//data,err:=ioutil.ReadFile("./static/view/index.html")
		//	if err!=nil{
		//		fmt.Println("ioutil.ReadFile error:%s\n",err.Error())
		//	}
		//
		//	io.WriteString(w,string(data))

	} else if r.Method == "POST" {

		/*2 r.FormFile 用于解析post 表单形式请求
		case:
		When you make a POST request, you have to encode the data that forms the body of the request in some way
		When you are writing client-side code:
		use multipart/form-data when your form includes any <input type="file"> elements*/
		/*	return :  multipart.File [io接口], *multipart.FileHeader【结构体】, error*/
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println("r.FormFile error:%s\n", err.Error())
		}

		//os.create 创建文件
		//in: filename
		//out: *file [可io],err
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Println("fail to create file:%s\n", err.Error())
		}

		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into file, err:%s\n", err.Error())
			return
		}

		// 3 重定向
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)

	}
}

// GetFileMetaHandler : 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {

	//默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，你才能对这个表单数据进行操作。

	//r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据

	//对文档不熟悉，可以自己检验
	//1	请求参数 ：http://localhost:8080/file/meta?filehash=47689294a3572f58fab8659516cecbb11227068b&filename="wang"
	//   被	r.ParseForm() 解析后 r.Form格式变成 ： map[filehash:[47689294a3572f58fab8659516cecbb11227068b] filename:["wang"]]
	//   请求参数 http://localhost:8080/file/meta?filehash=47689294a3572f58fab8659516cecbb11227068b&filename="wang"&filename="zhang"：
	//  map[filehash:[47689294a3572f58fab8659516cecbb11227068b] filename:["wang" "zhang"]]

	r.ParseForm()

	//r.Form里面包含了所有请求的参数，比如URL中query-string、POST的数据、PUT的数据
	//request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息

	//调用r.FormValue时会自动调用r.ParseForm，所以不必提前调用。r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
	//filehash := r.FormValue("filehash")
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)

	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//设置response的内容类型，位于响应头中【用类似结构体方法理解】
	//response响应头
	/*	是一系列key-value值  包含如下

		Cache-Control: private
		Content-Encoding: gzip
		Server: BWS/1.1
		Set-Cookie: delPer=0; path=/; domain=.baidu.com*/
	w.Header().Set("Content-type", "application/json")
	w.Write(data)

}
