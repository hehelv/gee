package gee

import (
	"fmt"
	"net/http"
)

//实现静态路由功能 , 基于http库实现

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

// New Engine构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 以下三个方法都是添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动服务
func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine) //engine要实现ServeHTTP方法，实现Handler接口
}

// 拦截所有的http请求到这个函数处理；
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 not found %s!", r.URL.Path)
	}
}
