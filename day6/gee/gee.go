package gee

import (
	"log"
	"net/http"
	"path"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

type (
	Engine struct {
		*RouterGroup //把engine抽象成最顶层的路由组
		router       *router
		groups       []*RouterGroup //存储所有的路由组
	}

	RouterGroup struct {
		prefix      string        //路由组通过前缀来划分
		middleWares []HandlerFunc //支持的中间件
		parent      *RouterGroup  //用于支持分组嵌套
		engine      *Engine       //通过engine 来使用engine的操作， 如addRoute
	}
)

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup} //将engine本身加入路由组中
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// Use 给路由组添加中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middleWares = append(group.middleWares, middlewares...)
}

// create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		// Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

// Static serve static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	// Register GET handlers
	group.GET(urlPattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middleWares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	engine.router.handle(c)
}
