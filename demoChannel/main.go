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

	linkCh := make(chan string)

	// 5 GET requests / start 5 go routines, immediately
	for _, link := range links {
		go checkLink(link, linkCh)
	}

	// // wait for 5 channel output
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println("LISTEN: ", <-linkCh, "replied")
	// }

	// loop true
	// wait for channel to return values
	// then assign to l
	// then run code in for body
	for l := range linkCh {
		// sleep() here will cause main routine
		// to wait before processing channel value
		// this is not ideal situation
		// time.Sleep(5 * time.Second)
		// go checkLink(l, linkCh)

		// solution:
		// waits 5 sec
		go func(link string) { // parameter name match
			time.Sleep(5 * time.Second)
			checkLink(link, linkCh) // parameter name match
		}(l) //pased in parameter (actual) value

	}

}

// GET request to the link
func checkLink(link string, c chan string) {
	// sleep() here not ideal
	//as it is expected this method to check the link immediately

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down! D:")
		c <- link
		return
	}

	fmt.Println(link, "is up! :D")
	c <- link
}
