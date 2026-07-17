//Write a merge function that takes exactly 3 input channels and merges them into 1 output channel. 
// Each input channel sends 5 values then closes. The merged output channel must close only after all 3 input channels are closed and drained.
//Constraint: use select with nil channel pattern for closing detection — not WaitGroup inside merge.

package main

import "fmt"

func merge(a,b,c <-chan int )<- chan int{
	output := make(chan int,15)
	go func(){
        for a!=nil || b!=nil || c!=nil {
			select{
			case v,ok := <- a:
				if(ok){
					output <- v
				}else{
					a = nil
				}
			case v,ok := <- b:
				if(ok){
					output <- v;
				}else{
					b = nil
				}
			case v,ok := <- c:
				if(ok){
					output <- v;
				}else{
					c= nil
				}
			
			}
		}
		close(output)
	}()
	return output
}

func main(){
    input1 := make(chan int,5)
	input2 := make(chan int,5)
	input3 := make(chan int,5)

	go func(){
		for i:=1;i<=5;i++{
			input1<-i
		}
		close(input1)
	}()
	go func(){
		for i:=6;i<=10;i++{
			input2<-i
		}
		close(input2)
	}()
	go func(){
		for i:=11;i<=15;i++{
			input3<-i
		}
		close(input3)
	}()

	output := merge(input1,input2,input3)

	for out := range output{
		fmt.Println(out)
	}




}