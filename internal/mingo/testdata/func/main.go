package main

import "fmt"

func hello() {
	func() {
		fmt.Println("Hello")
	}()

	func(s string) {
		fmt.Printf("Hello, %s\n", s)
	}("world")

	if err := func() error {
		return nil
	}(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	go hello()
}
