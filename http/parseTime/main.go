package main

import (
	"fmt"
	"net/http"
)

//ParseHTTPVersion parses an HTTP version string. "HTTP/1.0" returns (1, 0, true).

func main() {

	major, err := http.ParseTime("2020-10-08")
	if err != nil {
		fmt.Println("abc")
	}
	fmt.Println(major)
	//http.ListenAndServe(":8080", nil)
}
