package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	port := "8000"
	if len(os.Args) > 1 {
		port_i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Specified port values is incorrect", err)
			os.Exit(1)
		}
		if port_i <= 0 {
			fmt.Fprintf(os.Stderr, "Specified port values is incorrect: %s", os.Args[1])
			os.Exit(1)
		}
		port = os.Args[1]
	}
	fmt.Println("Starting server on http://0.0.0.0:" + port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
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
