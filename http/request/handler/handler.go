package handler

import "net/http"

func UploadFile(r *http.Request, w http.ResponseWriter) {
	// 1 r.Method 方法用来校验请求类型
	if r.Method == "GET" {

	} else if r.Method == "POST" {

		//When you make a POST request, you have to encode the data that forms the body of the request in some way
		//When you are writing client-side code:
		//use multipart/form-data when your form includes any <input type="file"> elements

		r.FormFile()
	}

}
