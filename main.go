package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://localhost:3000",
		"https://google.com",
		"https://Amazon.com",
		"https://golang.com",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checker(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checker(link string, c chan string) {
	reps, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "the server is down")
		c <- "The server is down"
		return
	}

	fmt.Println(link, "it working")
	fmt.Println(reps)
	c <- "Yep it working"
}
