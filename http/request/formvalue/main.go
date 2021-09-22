package main

import (
	"encoding/json"
	"net/http"
)

/*
   formvalue功能：
      1 : 输入 key , 返回请求路径后添加的参数键值对中的值
	如： http://localhost:9000/?file=abc
		formvalue(file) return "abc"

		其他
 		 if key is not present, FormValue returns the empty string.
*/

func main() {

	http.HandleFunc("/", testFormValue)
	http.ListenAndServe(":9000", nil)

}

func testFormValue(w http.ResponseWriter, r *http.Request) {
	res := r.FormValue("file")
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
