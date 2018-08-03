package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-timer.C:
			handler()
		}
	}
}

func handler() {
	fmt.Println(time.Now())
}