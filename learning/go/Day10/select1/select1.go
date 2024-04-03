package main
import (
	"fmt"
	"time"
)

func server1(ch chan <- string) {
	time.Sleep(2* time.Second)
	ch <- "One"
}

func server2(ch chan <- string) {
	time.Sleep(2 * time.Second)
	ch <- "two"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go server1(c1)
	go server2(c2)
	for i :=0; i< 2; i++ {
		select {
		case msg1:= <- c1:
			fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}

}