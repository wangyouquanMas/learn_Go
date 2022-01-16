package main

import "net/http"

// 功能： 获取表单数据

/*
	测试用例
 curl -d "name=wang&pass=1234" "localhost:8080/conclusion"
*/

func main() {
	http.HandleFunc("/conclusion", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		pass := r.FormValue("pass")
		w.Write([]byte(name + " " + pass))
	})

	http.ListenAndServe(":8080", nil)
}
