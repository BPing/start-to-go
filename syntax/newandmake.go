package syntax

import (
	"fmt"
	"github.com/start-to-go/syntax/interface"
)

func EchoNew() {
	p1 := new(int)
	fmt.Printf("p1 --> %#v \n ", p1)           //
	fmt.Printf("p1 point to --> %#v \n ", *p1) //0
	var p2 *int
	i := 0
	p2 = &i
	fmt.Printf("p2 --> %#v \n ", p2)           //
	fmt.Printf("p2 point to --> %#v \n ", *p2) //0

	pv := new(_interface.Vertex)
	fmt.Printf("pv --> %#v \n ", pv)           //
	fmt.Printf("pv point to --> %#v \n ", *pv) //0
}

func EchoMake() {
	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
	}
	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2) // []int{0, 0, 0}
	}

	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}
	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2)
		//map[int]string{}
	}

	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}
	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}
}
