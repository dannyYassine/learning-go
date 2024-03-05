package main

import (
	"fmt"
	"time"
)

func main() {
	var time time.Time = time.Now()

	switch {
	case time.Hour() < 12:
		fmt.Println("Good morning!")
	case time.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
