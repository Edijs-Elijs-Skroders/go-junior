package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// TODO: run the work function in a goroutine while in the for loop
// TODO: make sure all the workers finish 'working' before the WorkerFunc returns - you can add extra parameters to the work func
func WorkerFunc() {

	for i := 0; i < 10; i++ {
		work(i)
	}
}

func work(i int) {
	fmt.Printf("worker %v working \n", i)
	fmt.Println("sleeping")
	time.Sleep(time.Second * 2)
	fmt.Println("finished sleeping")

}


//TODO: create a function inside this this function that reads from the inputs channel using for-range loop
//TODO: the new function should take an int from the inputs, multiply it by 5, and put it in the res channel
//TODO: run the new function 5 times in a goroutine
//TODO:  in a loop 'i := 0; i < 100; i++', send the i to the inputs channel
//Tip: to make a function stop infinitely reading from a channel in for-range loop, close the channel using close(). It will still read all the messages that are left
//TODO: print out all the results using readResults function
func WorkerChannelFunc() {
	var wg sync.WaitGroup
	inputs := make(chan int)
	res := make(chan int, 100)

	
	f:= func(i int) {
		defer wg.Done()
		for v := range inputs {
			res <- v * 5
			fmt.Printf("worker %v working\n", i)
			time.Sleep(200 * time.Millisecond)
		}
	}

	for i:=0; i <5; i++ {
		wg.Add(1)
		go f(i)
	}

	for i := 0; i < 100; i++ {
		inputs <- i
	}

	close(inputs)
	wg.Wait()
	close(res)

	for r := range res {
		readResults(r)
	}

}


func readResults(i int) {
	fmt.Println(i)
}
