package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<-channel)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, channel chan<- string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}

	output := resp.Status
	channel <- fmt.Sprintf("%.2fs %s", time.Since(start).Seconds(), output)
}
