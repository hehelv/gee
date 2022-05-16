package main

/*
(1)
$ curl -i http://localhost:9090/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>
(2)
$ curl "http://localhost:9090/hello?name=jzc"
hello jzc, you're at /hello
(3)
$ curl "http://localhost:9090/hello/jzc"
hello jzc, you're at /hello/jzc
(4)
$ curl "http://localhost:9090/assets/css/jzc.css"
{"filepath":"css/jzc.css"}
(5)
$ curl "http://localhost:9090/xxx"
404 NOT FOUND: /xxx
*/

import (
	"net/http"

	"day3/gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9090")
}
