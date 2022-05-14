package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Get("/", indexHandler)

	r.Get("/hello", helloHandler)

	r.Run(":9090")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path: %q", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q] : %v\n", k, v)
	}
}
