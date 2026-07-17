//One generator sends 10 jobs (integers 1-10) into a jobs channel.
//Three worker goroutines all read from the same jobs channel and process each job (multiply by 2, sleep 100ms).
//All results go into one shared results channel. Main collects and prints all results.
//Constraint: exactly 3 workers, each job processed exactly once.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(job <- chan int, result chan <- int, wg *sync.WaitGroup){
	defer wg.Done()
	for val := range job{
		time.Sleep(100*time.Millisecond)
		result <- val*2
	}
}

func main(){
	job := make(chan int,10)
	result:=make(chan int,3)
	for i:=1;i<=10;i++{
		job <- i
	}
	close(job)
	var wg sync.WaitGroup
	for i:=0;i<3;i++{
        wg.Add(1)
		go worker(job,result,&wg)
	}
	go func(){
		wg.Wait()
		close(result)
	}()

	for res := range result{
        fmt.Println(res)
	}

}