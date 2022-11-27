package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

//实现了Handler接口的ServeHTTP方法
//之前我们只能针对具体路由写处理逻辑，现在可以拦截所有http请求，可以在这里加一些统一处理逻辑（日志、异常处理...）
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404ERROR\n")
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8081", engine))
}
