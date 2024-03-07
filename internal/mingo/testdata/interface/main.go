package main

type I1 interface {
	F1() string
	F2() (string, error)
	F3() (s string, err error)
}

type I2 interface {
	I1

	F4(s string)
	F5(string)
	F6(s1, s2 string)
	F7(s1 string, s2 string)
	F8(s1, s2 string, i int)
	F9(f func() string)
}
