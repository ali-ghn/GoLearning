package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// a slice of links to 5 most famous websites
	links := []string{
		"https://google.com",
		// "https://facebook.com",
		// Change facebook to duckduckgo 'cause it is Iran :)
		"https://duckduckgo.com",
		"https://amazon.com",
		// "https://twitter.com",
		// Change twitter to wikipedia 'cause it is Iran :)
		"https://wikipedia.org",
		// "https://golang.org",
		// Change golang to stackoverflow 'cause it is Iran :)
		"https://stackoverflow.com",
	}
	// loop through the links and make a request to see if they will respond with 200 OK
	// In this example my computer took about 50.589 seconds to complete this task (without concurrency)

	// for _, link := range links {
	// 	checkLink(link)
	// }

	// exact same thing but with concurrency (almost)
	// for _, link := range links {
	// 	go checkLink(link)
	// }

	// If you run the program you will see that the program would not show anything but why is that ?
	// the main function is actully a goroutin itself which holds the responsibility of running the actual program.
	// when creating child goroutins, the parent goroutine will not wait for the child goroutine to finish
	// therefore the application will get closed automatically and no results will be shown.
	// To fix this problem we can use Channel to communicate between the main goroutin and child goroutines.

	// create a channel to communicate between the main goroutine and child goroutines
	// make a channel of strings
	c := make(chan string)

	// loop through the links and make a request to see if they will respond with 200 OK
	for _, link := range links {
		// make a goroutine to check the link
		go checkLinkConcurrent(link, c)
	}

	// print the results of the channel once

	// fmt.Println(<-c)

	// If you run the program you will see that whatever goroutine is sending back the data to channel will be printed first
	// and all the other gourotines will be terminated as the main function only waits for the first goroutine to finish.
	// To fix this problem we can use select to wait for the data to be sent back to the channel.

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// using this method my computer took about 0.13.976 seconds to complete this task (with concurrency)

	// You can also run the program for infinite loop and see the results.
	// for {
	// 	go checkLinkConcurrent(<-c, c)
	// }

	// Or a nicer format of creating such a loop
	// for l := range c {
	// 	go checkLinkConcurrent(l, c)
	// }

	// Well nice but we're just making requests to the websites over and over again without any delays in between.
	// we can fix that though using functions literals.

	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLinkConcurrent(link, c)
		}(l)
	}

}

// checkLink checks if a link returns a 200 response
func checkLink(link string) error {
	// make a request to the link
	r, err := http.Get(link)
	if err != nil {
		return err
	}
	// close the response body after we are done with it
	defer r.Body.Close()
	// check if the response status code is 200 OK
	if r.StatusCode == 200 {
		fmt.Println(link, ": Status OK")
	} else {
		return fmt.Errorf("%s: %s", link, r.Status)
	}
	return nil
}

// checkLinkConcurrent checks if a link returns a 200 response
func checkLinkConcurrent(link string, c chan string) {
	// make a request to the link
	r, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- link
		return
	}
	// close the response body after we are done with it
	defer r.Body.Close()
	// check if the response status code is 200 OK
	if r.StatusCode == 200 {
		fmt.Println(link, ": Status OK")
		c <- link
	} else {
		fmt.Println(link, ":", r.Status)
		c <- link
	}
}
