package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
var port = flag.String("p", "8000", "port")

func main() {
	flag.Parse()
	fmt.Println("Starting server on http://0.0.0.0:" + *port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+*port, nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/count":
		fmt.Fprintf(writer, "Count: %d\n", count)
	case "/favicon.ico":
	default:
		{
			mu.Lock()
			count++
			fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
			mu.Unlock()
		}
	}
}
