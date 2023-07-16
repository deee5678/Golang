package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { //watch the channel c & whenever a value comes out of it,assign the value to l & body of the for loop is immedeiately executed then
		go func(link string) { // function literal [ananymous functn]
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // calling the function literal

		/*time.Sleep(5 * time.Second)
		go checkLink(l, c)*/
	}
	/*for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}*/
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // c <- "Might be down i think"
		//return
	}
	fmt.Println(link, "is up!")
	c <- link // c <- "yep it's up"
}
