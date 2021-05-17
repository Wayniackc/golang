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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Infinite loop
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	// channel <- 5 - Send the value '5' into this channel
	// myNumber <- channel - Wait for a value to be sent into the channel. When we get one, assign the value to 'myNumber'
	// fmt.Println(<- channel) - Wait for a value to be sent into the channel. When we get one, log it out immediately.

	fmt.Println(link, "is up!")
	c <- link
}
