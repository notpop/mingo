package main

const (
	A           = 1
	B       int = 2
	C, D    int = 3, 4
	E, F, G     = 3, 4, 5
)

const (
	_ = iota
	a
	b
	c
)

const w = 1
const x int = 2
const y, z = 3, 4

func main() {
	const a = 1
	const b, c = 2, 3
	const (
		d = 4
		e = 5
	)
}
