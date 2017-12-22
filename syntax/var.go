package syntax

import (
	"fmt"
)

const Pi = 3.14159

var (
	aPtr   *int
	strPtr *string
)

func Echo() {
	var a int = 15
	var i = 5
	var b bool = false
	var str string = "Go says hello to the world!"

	fmt.Println(a, i, b, str)

	str1 := `:=`
	fmt.Println(str1)

	aPtr = &a
	strPtr = &str
	*aPtr = *aPtr + 1
	*strPtr = "pointer str"
	fmt.Println(a, str)

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
}
