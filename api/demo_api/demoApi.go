package demo_api

import (
	"fmt"
	"net/http"
)

func Demo1(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "demo1")
}

func Demo2(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "demo2")
}
