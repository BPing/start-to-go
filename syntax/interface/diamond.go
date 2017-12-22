package _interface

import "fmt"

type A struct {
	name string
}

func (A) Name() string {
	return "A"
}

type B struct {
	A
}

func (B) Name() string {
	return "B"
}

type C struct {
	A
}

func (C) Name() string {
	return "C"
}

type D struct {
	B
	C
}

func EchoDiamond() {
	//d := new(D)
	//fmt.Println(d.Name())
	fmt.Println("dd")
}
