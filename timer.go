package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTicker(3 * time.Second)

	ch := make(chan time.Time, 10)

	go func() {
		for {
			select {
			case <-timer.C:
				ch <-time.Now()
				handler()
			}
		}
	}()

	<-time.After(time.Second * 20)

	for t := range ch {
		fmt.Println(t)
	}

	close(ch)
}

func handler() {
	fmt.Println(time.Now())
}