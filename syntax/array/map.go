package array

import (
	"fmt"
)

//
// https://blog.go-zh.org/go-maps-in-action
// go1.9 sync.Map
//

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func EchoMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	//
	// map 的文法
	//
	//var m = map[string]Vertex{
	//	"Bell Labs": Vertex{
	//		40.68433, -74.39967,
	//	},
	//	"Google": Vertex{
	//		37.42202, -122.08408,
	//	},
	//}

	// var m = map[string]Vertex{
	//	"Bell Labs": {40.68433, -74.39967},
	//	"Google":    {37.42202, -122.08408},
	// }

	fmt.Println(m["Bell Labs"])
}

func ModifyMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func RangeMap() {
	// 遍历随机性
	m := map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
		"Bai du": Vertex{
			25.42202, -142.08408,
		},
	}
	fmt.Println("ONE----")
	for k, v := range m {
		fmt.Println("The value:", v, "The key:", k)
	}
	fmt.Println("TWO----")
	for k, v := range m {
		fmt.Println("The value:", v, "The key:", k)
	}
}
