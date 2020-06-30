package main

import "github.com/ritabc/gotr/019-reflection/part-1/ex5/foo"

type (
	ID     string
	Person struct {
		name string
	}
)

// Show that fmt.Println() takes any interface
func main() {
	foo.Println(true)
	foo.Println(2010)
	foo.Println(9.15)
	foo.Println(7 * 7i)
	foo.Println("Hello World!")
	foo.Println(ID("19438572"))
	foo.Println([5]byte{})
	foo.Println([]byte{})
	foo.Println(map[string]int{})
	foo.Println(Person{name: "Jane Doe"})
	foo.Println(&Person{name: "Jane Doe"})
	foo.Println(make(chan int))
	foo.Println(nil)
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
