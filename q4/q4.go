//E4 — WaitGroup Basic
//Launch 5 goroutines. Each prints its ID and sleeps a random duration (0-500ms). 
// Main must wait for all 5 to finish before printing "all done". Use WaitGroup only — no channels.


package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

func worker(id int, wg *sync.WaitGroup){
	fmt.Println("your id is=",id)
	t := rand.Intn(500)
	time.Sleep(time.Duration(t)*time.Millisecond)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	for i:=1;i<=5;i++{
        wg.Add(1)
		go worker(i,&wg)
	}
	wg.Wait()
}