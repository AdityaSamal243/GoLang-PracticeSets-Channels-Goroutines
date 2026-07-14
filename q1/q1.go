//Two goroutines. One sends 5 integers one by one. 
// Main receives and prints each. No WaitGroup allowed — use channel completion signal only.


package main	

import (
	"fmt"
)

func main(){
	// done := make(chan struct{})
	value := make(chan int,1)

	go func(){
		for i:=0;i<5;i++{
			value <- i
		}
		close(value)
	}()

	for val := range value{
		fmt.Println(val)
	}

}