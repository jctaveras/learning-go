package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	for url := range c {
		go func(chanURL string) {
			time.Sleep(time.Second * 5)
			checkURL(chanURL, c)
		}(url)
	}
}

func checkURL(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println("Server: ", url, " might be down")
		c <- url
		return
	}

	fmt.Println("Server: ", url, " is running")
	c <- url
}
