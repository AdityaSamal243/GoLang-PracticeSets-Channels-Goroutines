//Create a buffered channel of size 5. Send 5 jobs (strings) into it from main without blocking.
// Launch one worker goroutine that drains it and prints each job with a 100ms delay between each.

package main

import (
	"fmt"
	"time"
)

func main() {
	value := make(chan string, 5)
	done := make(chan struct{})

	go func() {
		for val := range value{
			fmt.Println(val)
			time.Sleep(1*time.Second)
		}
		done <- struct{}{}

	}()

	names := []string{"aditya", "tanu", "adi", "tan", "work"}
	for _, name := range names {
		value <- name
	}
	close(value)
	<- done // used this so it doesn't terminate after sending the values .. wait to recieve the signal 

}
