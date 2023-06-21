package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 处理函数
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.path=%q\n", request.URL.Path)
}

//
