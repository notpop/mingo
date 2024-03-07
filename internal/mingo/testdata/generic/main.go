package main

import "fmt"

func Ptr[T any](p T) *T {
	return &p
}

func Equals[T, U comparable](t T, u U) bool {
	return t == u
}

func Hoge[T comparable, U fmt.Stringer](s []T, e U) bool {
	return true
}

func main() {
	fmt.Println(Ptr(1))
}
