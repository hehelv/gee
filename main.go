package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	http.ListenAndServe(":9090", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "path: %q", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	for k, v := range r.Header {
		fmt.Fprintf(w, "header[%q] : %v\n", k, v)
	}
}


