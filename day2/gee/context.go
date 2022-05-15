package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 为添加json格式数据方便
type H map[string]interface{}

// Context 设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。
// Context 返回数据，设置中间件，动态路由等功能
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	Path   string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

// PostForm 通过key获得表单数据
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 通过url获取Query参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置http状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 主要用来设置返回数据类型
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// 下面四个都是回调，返回数据
func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Json(code int, obj interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
