//H1 — Pipeline With Backpressure
// Build a pipeline where the producer generates 100 items as fast as possible.
// The consumer processes each item with a 50ms delay.
// Add backpressure — the producer must slow down when the consumer can't keep up.
// Prove backpressure is working by printing producer and consumer rates every second.
// Producer rate should drop to match consumer rate after initial burst.

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var prodCount int64
var consCount int64

func generator(Initial chan <- int){
	fmt.Println("started generating")
	for i:=1;i<=100;i++{
        Initial <- i;  // backpressure point-- once buffer is full it stops here 
		atomic.AddInt64(&prodCount,1)
	}
	close(Initial)
}
func consumer(Initial <- chan int,done chan <- struct{}){
	fmt.Println("started consuming")
	for range Initial{
		time.Sleep(50*time.Millisecond)
		// fmt.Println(val)
		atomic.AddInt64(&consCount,1)
	}
	done <- struct{}{}
}

func main(){
	Initial := make(chan int,3)
	done := make(chan struct{})
	go generator(Initial)
	go consumer(Initial,done)
    
	monitor := make(chan struct{})

	go func(){
		ticker := time.NewTicker(1*time.Second)
		defer ticker.Stop()

		fmt.Printf("\n%-15s | %-20s | %-20s\n", "Elapsed Time", "Producer Rate/sec", "Consumer Rate/sec")
		fmt.Println("------------------------------------------------------------------")
        start := time.Now()
		for{
			select{
			case <- monitor:
				return
			case <-ticker.C:
				pRate := atomic.SwapInt64(&prodCount,0) // return old value and swap prodcount with 0
				cRate := atomic.SwapInt64(&consCount,0)
                fmt.Printf("%-15s | %-20d | %-20d\n", 
					time.Since(start).Round(time.Second), pRate, cRate)
			}
		}

	}()

	<-done
	close(monitor)
	fmt.Println("all finished")

}  