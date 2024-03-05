package main

import "fmt"

func main() {
	defer fmt.Println("Exited...")

	fmt.Println("Main function")
}
