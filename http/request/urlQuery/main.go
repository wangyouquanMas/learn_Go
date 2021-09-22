package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	req.url.query 作用

	1 ：解析url 将参数设置为键值对返回
	例 ： http://localhost:9000/?NAME=%22wang%22&ange=25
	    返回 map[NAME:["wang"] ange:[25]]
*/

func main() {

	http.HandleFunc("/", testFormValue)
	http.ListenAndServe(":9000", nil)

}

func testFormValue(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query()
	fmt.Println(res)
	//obj := map[string]interface{}{
	//	"name":     "geektutu",
	//	"password": "1234",
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
