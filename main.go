package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// WebsiteStatus is great
type WebsiteStatus struct {
	url    string
	status bool
	err    error
	time   time.Duration
}

func main() {
	urls := []string{
		"http://facebook.com",
		"http://amazon.com",
		"http://syhol.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.com",
	}

	courseWay(urls)
}

func checkURL(url string) WebsiteStatus {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)
	return WebsiteStatus{
		url:    url,
		status: err == nil && resp.StatusCode >= 200 && resp.StatusCode < 400,
		err:    err,
		time:   duration,
	}
}

func reportOnURLStatus(ws WebsiteStatus) {
	if ws.err != nil {
		log.Println(ws.err)
	}
	if ws.status {
		fmt.Println(ws.url, "is up, finnished in", ws.time)
	} else {
		fmt.Println(ws.url, "is down")
	}
}

////////////
// My Way //
////////////

func myWay(urls []string) {
	c := make(chan WebsiteStatus)

	for _, url := range urls {
		go asyncInfiniteCheckURL(url, c)
	}

	for {
		reportOnURLStatus(<-c)
	}
}

func asyncInfiniteCheckURL(url string, c chan WebsiteStatus) {
	for {
		c <- checkURL(url)
		time.Sleep(3 * time.Second)
	}
}

////////////////
// Course Way //
////////////////

func courseWay(urls []string) {
	c := make(chan string)

	for _, url := range urls {
		go asyncRecursiveCheckURL(url, c)
	}

	for url := range c {
		go func(url string) {
			time.Sleep(3 * time.Second)
			asyncRecursiveCheckURL(url, c)
		}(url)
	}
}

func asyncRecursiveCheckURL(url string, c chan string) {
	reportOnURLStatus(checkURL(url))
	c <- url
}
