package main

import (
	"fmt"
	"time"
)

func main() {
	requests:= make(chan int, 5)
	for i := 1; i<=5; i++{
		requests <- i
	}
	close(requests)
	limiter := time.Tick(1 * time.Second)

	// rate limited, display at a 1 second interval
	for reqId := range requests {
		<- limiter
		fmt.Println("Limited Request", reqId, time.Now())
	}
	burstyLimiter := make(chan time.Time ,3)

	for i:=1;i<=3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(1* time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i:= 1;i<=5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Bursty Request", req, time.Now())
	}
}