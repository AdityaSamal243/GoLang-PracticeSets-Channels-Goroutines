//M6 — Or-Done Channel
//Write an orDone function:
//gofunc orDone(done, c <-chan int) <-chan int
//It reads from c but stops when done is closed. 
// The returned channel closes when either c closes or done closes. Demonstrate it by stopping a long-running stream after 3 values.

package main	

import (
	"fmt"
)

func worker(done chan struct{}, c <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case result <- v:
				case <-done:
					return
				}
			}
		}
	}()
	return result
}

func main(){
	c := make(chan int)
	done := make(chan struct{})
	
	go func(){
		defer close(c)
		for i:=1; ;i++{
			select{
			case <- done:
				fmt.Println("procucer stopped")
				return
			case c <- i:
			}
		}
	}()

	count := 0
	values := worker(done,c)

	for val := range values{
		fmt.Println(val)
		count++

		if count == 3{
			close(done)
			break
		}
	}

	fmt.Println("main exited")


}