package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string, c chan string) {
	println(link, ".......⏳")
	_, err := http.Get(link)
	if err != nil {
		println(link, "might be down ❌")
		c <- "Might be down I think"
		return
	}
	println(link, "is up ✅")
	c <- "Yep it is up"
}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}
