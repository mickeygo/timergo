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
			fmt.Println(time.Now())
		}
	}	
}