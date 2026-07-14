
// Write a pipeline with three functions:

// generate(nums ...int) <-chan int — sends ints into a channel
// square(in <-chan int) <-chan int — receives, squares each, sends out
// main — receives from square and prints

// All channel parameters must be directional (<-chan or chan<-). No bidirectional channels anywhere except creation.


package main	

import(
	"fmt"
)

func generate(nums...int) <- chan int{
	value := make(chan int,5)
	go func(){
		for _,num := range nums{
			value <- num
		}
		close(value)
	}()
	return value
}

func squarer(value <-chan int) <-chan int {
	result := make(chan int, 5)
	go func(){
		for val := range value {
			x := val * val
			result <- x
		}
		close(result)
	}()
	return result
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	// value := make(chan int,5)
	// result := make(chan int,5)
	generator := generate(nums...)
	result :=  squarer(generator)

	for res := range result{
		fmt.Println(res)
	}

}