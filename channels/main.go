package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// channel
	c := make(chan string)

	for _, link := range links {
		// go keywordはfunctionの前だけ
		go checkLink(link, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }
	// 上と同じ
	for l := range c {
		// time.Sleep(2 * time.Second)
		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
