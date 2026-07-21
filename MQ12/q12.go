//Launch 5 goroutines. Each does some work that might fail (simulate: goroutine 3 always returns an error). 
// Collect ALL errors (not just the first).
// Main prints each error and a final count of how many goroutines failed. 
// No external packages — implement error collection manually using channels.

package main

import (
	"fmt"
)

func main(){
	// create an error channel of size 5 since 5 goroutines.
	const n = 5
	errCh := make(chan error, n)
    // simulate 5 goroutines 
	for i := 0; i < n; i++ {
		go worker(i, errCh)
	}
    // run a loop through errch channel and capture all erros in err slice. 
	var errs []error
	for i := 0; i < n; i++ {
		if err := <-errCh; err != nil {
			errs = append(errs, err)
		}
	}
    // iterate through slice and print all errors
	for _, e := range errs {
		fmt.Println("error:", e)
	}
	fmt.Printf("%d goroutines failed\n", len(errs))
}


func worker(id int, errCh chan<- error) {
	// simulate work
	if id == 3 {
		errCh <- fmt.Errorf("worker %d failed", id)
		return
	}
	// succeed
	errCh <- nil
}
