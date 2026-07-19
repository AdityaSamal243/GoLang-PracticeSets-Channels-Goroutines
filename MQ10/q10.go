//You have 20 URLs to "fetch" (simulate with random sleep 100-500ms).
// Maximum 3 fetches can happen concurrently at any time. Use a buffered channel as a counting semaphore. Print when each fetch starts and completes.
// Total time should be significantly less than sequential but fetch count in-flight should never exceed 3.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetcher(url int, sem chan struct{}, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	//aquire semaphore lock
	sem <- struct{}{}
	defer func(){
		<-sem  
	}()
	fmt.Println("started job=", url)
	wait := rand.Intn(400) + 100
	time.Sleep(time.Duration(wait) * time.Millisecond)
	fmt.Println("finished job=", url)
	result <- url 
}
func main() {
	sem := make(chan struct{}, 3)
	result := make(chan int, 20)
	// simulatiing numbers 1 to 20 as url for writing ease
	var wg sync.WaitGroup
	for i := 1; i < 20; i++ {
		wg.Add(1)
		go fetcher(i, sem, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println("collected", res)
	}
}
