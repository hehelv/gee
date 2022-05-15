package gee

import "net/http"

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

// New 构造函数
func New() *Engine {
	return &Engine{router: NewRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

// 实现Handler接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 创建一个context对象， context对象从请求开始产生，到服务器结束消失
	c := newContext(w, r)
	engine.router.handle(c)
}
