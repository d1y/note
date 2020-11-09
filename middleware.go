package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// `cors` 中间件
// func useCORS(r *httprouter.Router) {
// 	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var path = r.URL.Path
// 		fmt.Println("path: ", path)
// 		if r.Header.Get("Access-Control-Request-Method") != "" {
// 			// Set CORS headers
// 			header := w.Header()
// 			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
// 			header.Set("Access-Control-Allow-Origin", "*")
// 		}

// 		// Adjust status code to 204
// 		w.WriteHeader(http.StatusNoContent)
// 	})
// }

// 静态文件
func useStatic(r *httprouter.Router) {
	r.ServeFiles("/static/*filepath", http.Dir("./static"))
}
