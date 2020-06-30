package foo

import "fmt"

// Type switching is not great because it can't handle custom types from different packages.
func Println(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Println(x.(bool))
	case int:
		fmt.Println(x.(int))
	case float64:
		fmt.Println(x.(float64))
	case complex128:
		fmt.Println(x.(complex128))
	case string:
		fmt.Println(x.(string))
	case chan int:
		fmt.Println(x.(chan int))
	default:
		fmt.Println("Unknown type")
	}
}
