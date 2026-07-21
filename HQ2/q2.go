//H2 — Concurrent Ordered Results
//Launch 10 goroutines. Each takes a job ID and returns Result{ID, Value} after a random delay.
// Main must print results in job ID order (1, 2, 3... 10) even though they complete out of order.
// Two approaches exist — implement the more efficient one (hint: not collect-all-then-sort).

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct{
	id int
	value int
}

func worker(id int,Pipe chan<-Result, wg *sync.WaitGroup){
	   defer wg.Done()
       value := rand.Intn(50)
	   delay := rand.Intn(400)+100
	   time.Sleep(time.Duration(delay)*time.Millisecond)
	   Pipe <- Result{id:id , value:value}
}


func main(){
	Pipe := make(chan Result,10)
	var wg sync.WaitGroup
	for i:=1;i<=10;i++{
		wg.Add(1)
		go worker(i,Pipe,&wg)
	}
	go func(){
		wg.Wait()
		close(Pipe)
	}()

	buffer := make(map[int]int)
	nextId := 1

	for val:= range Pipe{
        buffer[val.id]= val.value
		for{
			if body,found := buffer[nextId]; found{
				fmt.Println(nextId,":",body)
				delete(buffer,nextId)
				nextId++
			}else{
				break;
			}
		}
	}
}