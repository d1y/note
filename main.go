package main

import (
	"fmt"
	"net/http"
)

func h(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/h", h)
	http.HandleFunc("/api", headers)

	http.ListenAndServe(":2333", nil)
}
