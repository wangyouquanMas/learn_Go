package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/conclusion", func(writer http.ResponseWriter, request *http.Request) {
		obj := map[string]interface{}{
			"name":     "dong",
			"password": "yang",
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(writer)
		if err := encoder.Encode(obj); err != nil {
			http.Error(writer, err.Error(), 500)
		}
	})
	http.ListenAndServe(":8000", nil)
}
