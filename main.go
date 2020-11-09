// create by d1y<chenhonzhou@gmail.com>

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/d1y/note/api"
	v1 "github.com/d1y/note/api/v1"
	"github.com/d1y/note/conf"
	_ "github.com/d1y/note/db"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var app = httprouter.New()

	// 跨域
	// useCORS(app)

	// 静态文件
	useStatic(app)

	var htmlTemplate, err = template.ParseFiles("./template/index.gohtml")
	if err != nil {
		panic(err)
	}

	// ========

	app.GET("/", api.Index302)

	var webMatchPath = fmt.Sprintf("/%s/:path", conf.WebPrefix)

	app.GET(webMatchPath, func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		var path = p.ByName("path")
		var o = v1.GetData(path)
		htmlTemplate.Execute(w, map[string]string{
			"title":   path,
			"content": o.Data.Content,
		})
	})

	app.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.Index302(w, r, nil)
	})

	// ========

	// ========

	app.GET("/api/router/:router", api.GetRouterData)

	app.POST("/api/router", api.PostRouterData)

	// ========

	runHTTP(app)
}

func runHTTP(app *httprouter.Router) {
	var listener = ":" + strconv.Itoa(conf.ExposePort)
	var msg = fmt.Sprintf("http://localhost:%v", conf.ExposePort)
	fmt.Println("http server listener to: ", msg)
	log.Fatal(http.ListenAndServe(listener, app))
}
