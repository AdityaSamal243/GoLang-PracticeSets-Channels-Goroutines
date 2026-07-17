//Build a 3-stage pipeline: generate → process → collect. Add a done channel.
// If done is closed at any stage, that stage exits cleanly without goroutine leaks.
// Test by closing done after 3 items have been processed —
// verify remaining items are not processed and no goroutines leak (use runtime.NumGoroutine() before and after).

package main

import (
	// "runtime"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done chan struct{}) <-chan int {
	process := make(chan int)
	go func() {
		defer fmt.Println("generator exited")
		defer close(process)
		for i := 1; i <= 6; i++ {
			select {
			case <-done:
				return
			case process <- i:
			}
		}
	}()
	return process
}
func processor(processed <-chan int, done chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer fmt.Println("processor exited")
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v, ok := <-processed:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case result <- v * 2:
				}
			}
		}
	}()
	return result
}

func collector(results <-chan int, done chan struct{}) {
	count := 0
	for res := range results {
		fmt.Println(res)
		count++
		if count == 3 {
			close(done)
			return
		}

	}
}

func main() {
	var wg sync.WaitGroup
	before := runtime.NumGoroutine()
	done := make(chan struct{})
	processed := generator(done)
	results := processor(processed, done)
	wg.Add(1)
	go func() {
		defer wg.Done()
		collector(results, done)
	}()
	wg.Wait()
	time.Sleep(10*time.Millisecond)
	after := runtime.NumGoroutine()

	fmt.Println("initial goroutine=", before)
	fmt.Println("final goroutines=", after)

}
