package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	switch os {
	case "linux":
		fmt.Println("Linux is awesome")
		break
	default:
		fmt.Println("OS not found")
	}
}
