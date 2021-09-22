package main

import (
	"net/http"
)

/*
	httpResponseWriter.writerHeader作用

	1 ：给响应头的 status code 键赋值
*/

func main() {

	http.HandleFunc("/", testFormValue)
	http.ListenAndServe(":9000", nil)

}

func testFormValue(w http.ResponseWriter, r *http.Request) {

	//w.WriteHeader(200)
	//obj := map[string]interface{}{
	//	"name":     "geektutu",
	//	"password": "1234",
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	//encoder := json.NewEncoder(w)
	//if err := encoder.Encode(res); err != nil {
	//	http.Error(w, err.Error(), 500)
	//}
}
