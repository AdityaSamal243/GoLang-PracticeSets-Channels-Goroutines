// Worker Pool With Dynamic Scaling
//Build a worker pool that:

//Starts with 2 workers
// Adds a worker when job queue length exceeds 5
// Removes a worker when queue has been empty for 2 seconds
// Maximum 5 workers, minimum 1
// Print current worker count whenever it changes
// Test with a burst of 20 jobs followed by silence.

package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

func Worker(id int, jobs <-chan int, quit <- chan struct{},wg *sync.WaitGroup){
	defer wg.Done()
	for{
		select{
		case <-quit:
			fmt.Printf("worker %d removed - \n",id)
			return
		case v,ok := <- jobs:
			if !ok{
				return
			}
			time.Sleep(300*time.Millisecond)
			fmt.Printf("worker %d processed job %d \n",id,v)
		}
	}
}

//pool manager — owns count, owned by ONE goroutine only
// communicates entirely through channels

func manage(jobs <- chan int,wg *sync.WaitGroup,done <-chan struct{}){
	ticker := time.NewTicker(100*time.Millisecond)
	defer ticker.Stop()

	count:=0
	idleSince := time.Now()
	//TRACK QUIT CHANNELS- 1 per worker
	quits := make([]chan struct{},0,5)

	add := func(){
		if count >=5{
			return
		}
		quit := make(chan struct{})
		quits = append(quits,quit)
		count++
		fmt.Println("Worker added- pool size=",count)
		wg.Add(1)
		go Worker(count,jobs,quit,wg)
	}
	remove := func(){
		if count <=1 {
			return
		}
		last :=len(quits) - 1
		close(quits[last])
		quits = quits[:last]
		count--
		fmt.Println("Worker removed - pool size=",count)
	}
	add()
	add()

	for{
		select{
		case <- done:
			//shut down all worker
			for _,q := range quits{
                 close(q)
			}
			return
		case <-ticker.C:
			qlen := len(jobs)
			if qlen > 0{
				idleSince = time.Now()
			}
			switch{
			case qlen > 5 && count < 5:
				add()
			case qlen == 0 && time.Since(idleSince) > 2*time.Second && count > 1:
				remove()
			}
	
		}
	}
}


func main(){
	jobs := make(chan int,20)
	done := make(chan struct{})
	var wg sync.WaitGroup

	go manage(jobs, &wg, done)

	//20 jobs
	go func(){
		for i:= 1;i<=20;i++{
           jobs <- i 
		}
		fmt.Println("all jobs sent")
	}()
	
	time.Sleep(12 * time.Second)
	close(done)
	close(jobs)
	wg.Wait()
	fmt.Println("pool shut down cleanly")

}