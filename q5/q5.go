//Launch a goroutine that prints numbers 1 to infinity in a loop with 200ms sleep. 
// Main should stop it after 1 second using a done channel. No context, no time.After inside the goroutine.

package main

import (
	"fmt"
	"time"
)

func main(){
	done := make(chan struct{})

	go func(){
		i := 1
		for {
			select {
			case <-done:
				fmt.Println("shutting down at:",i)
				return
			default:
				fmt.Println(i)
				time.Sleep(200 * time.Millisecond)
				i++
			}
		}
	}()

	time.Sleep(1*time.Second)
	done <- struct{}{}
	time.Sleep(200*time.Millisecond)
	fmt.Println("closed cleanly")
}