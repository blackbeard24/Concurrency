package main

import (
	"fmt"
	"time"
)

const waitPeriod = 5 * time.Second

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	fmt.Println("Learning Concurrency")

	processes := map[string]int{"process1": 20, "process2": 18, "process3": 27}
	for process, completionTime := range processes {
		go operation(process, completionTime)
	}
	time.Sleep(waitPeriod)
}

func operation(process string, completionTime int) {
	for completionTime > 0 {
		fmt.Printf("%s still processing, time left %d \n", process, completionTime)
		completionTime = completionTime - 5

	}
	if completionTime <= 0 {
		fmt.Printf("%s completed", process)
	}
	time.Sleep(waitPeriod)
}
