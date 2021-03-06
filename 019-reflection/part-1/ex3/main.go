package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

// Show that fmt.Println() takes any interface
func main() {
	Println(true)
	Println(2010)
	Println(9.15)
	Println(7 * 7i)
	Println("Hello World!")
	Println(ID("19438572"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "Jane Doe"})
	Println(&Person{name: "Jane Doe"})
	Println(make(chan int))
	Println(nil)
}

func Println(x interface{}) {
	// x.type - for type informayion
	// x.value - for the value assigned

	// Use type assertion
	if v, ok := x.(ID); ok {
		fmt.Printf("x has type ID, which I defined. The value is: %v\n", v)
	} else {
		fmt.Printf("'%T' is not the type I want\n", x)
	}
}

/* When assigning 'var foo interface{}', two hidden fields are created: Value & Type.
/// When writing foo = 3.14, Value = 3.14 && Type = float64
/// When creating person literal and taking address of it, and writing foo = &Person{}, Value = address && Type = *Person
*/

/* Dynamic vs. Static Type
What is 'foo's type?
What if we, after saying foo = 3.14, we say: var goo = foo. What is 'goo' type?
What if we, after saying foo = &Person{}, we say hoo = foo. What is 'hoo' type?
The static type of foo is empty interface: interface{}. Go is statically typed language. Upon declaration, we assign the static type
It appears the Type of foo is changing, but it's the dynamic type that is changing, or the word/hidden field Type: float64 || Type: *Person that is changing
*/

// Type assertion definition: Mechanism to reveal dynamic type. Static type of foo is always empty interface
// Type switching: usint type as switch case condition
