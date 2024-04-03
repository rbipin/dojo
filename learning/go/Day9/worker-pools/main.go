package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("Starting Worker ", id, "Starting Job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "Completed Job", j)
		result <- j
	}
}



func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		fmt.Println("Start Worker ", w)
		go worker(w, jobs, result)
	}

	for i:=1;i<=5; i++ {
		jobs <- i
	}
	close(jobs)
	for a := 1; a<=numJobs; a++ {
		<- result
	}

}
