//M3 — Select With Timeout Per Operation
//A function fetchWithTimeout(url string, timeout time.Duration) (string, error) simulates fetching a URL.
// Internally it launches a goroutine that sleeps a random 0-2 seconds then sends a fake response.
//If the goroutine doesn't respond within timeout, return an error. Call this function 5 times concurrently
// and print each result as it arrives.

package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
func fetchWithTimeout(url string, timeout time.Duration)(string , error){
	 out := make(chan string,1)
	 
	 // simulating fetching url
     go func(){
	   ra := rand.Intn(3)
       time.Sleep(time.Duration(ra)*time.Second)
       out <- url
	 }()
     select{
	 case res := <-out:
		return res,nil
	 case <- time.After(timeout):
		return "", fmt.Errorf("timeout fetching %s", url)
	 }
}

func main(){
	urls := []string{"hello.com","world.com","good.com","foo.com","bar.net"}
    var wg sync.WaitGroup
	for _,url := range urls{
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			
			res,err := fetchWithTimeout(url,1*time.Second)
			if err!=nil{
				fmt.Println(err)
				return
			}
			fmt.Println(res)
		}(url)
	}
	wg.Wait()
}