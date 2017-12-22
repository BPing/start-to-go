package _interface

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X float64
	Y float64
}

func EchoStruct() {
	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略
		v3 = Vertex{}      // X:0 和 Y:0
		p  = &Vertex{1, 2} // 类型为 *Vertex
	)
	fmt.Println(Vertex{1, 2})

	fmt.Println(v1, v2, v3, p)
}

type TagType struct {
	// tags
	field1 bool   `tag:"An important answer"`
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func EchoStructTag() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

func EchoStructSelf() {
	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
}
