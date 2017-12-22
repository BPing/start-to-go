package syntax

import "fmt"

func For1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}

func ForRange() {
	arr := [...]string{"a", "b", "c"}
	for i := range arr {
		fmt.Println("Array item", i, "is", arr[i])
	}

	mapd := map[string]string{
		"a": "av",
		"b": "bv",
		"c": "cv",
	}
	for k, v := range mapd {
		fmt.Println("key :", k, " value:", v)
	}

	//
	//
}
