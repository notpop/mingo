package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello, 世界")
	time.Sleep(1 * time.Second)
}

func main() {
	go func() {
		fmt.Println("Hello, 世界")
		time.Sleep(1 * time.Second)
	}()

	go hello()
}
