package main

import "fmt"

func main() {
	fmt.Println("A")
	goto label
	fmt.Println("B")
label:
	fmt.Println("C")
}
