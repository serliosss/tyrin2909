package main

import "fmt"

// 29.09
// ЗАДАЧА №1
func main() {
	a1 := 10
	var a2 int = 11
	var a3 = 12

	b1 := "first"
	var b2 string = "second"
	var b3 = "third"

	c4 := 1.44
	var c5 float64 = 1.55
	var c6 = 1.66

	d1 := true
	var d2 bool = true
	var d3 = true

	e1 := "1"
	var e2 rune = '1'
	var e3 = "1"

	f1 := 'Z'
	var f2 byte = 'Z'
	var f3 = 'Z'

	fmt.Printf("INT\n")
	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%T\n\n", a3)

	fmt.Printf("STRING\n")
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)
	fmt.Printf("%T\n\n", b3)

	fmt.Printf("FLOAT\n")
	fmt.Printf("%T\n", c4)
	fmt.Printf("%T\n", c5)
	fmt.Printf("%T\n\n", c6)

	fmt.Printf("BOOL\n")
	fmt.Printf("%T\n", d1)
	fmt.Printf("%T\n", d2)
	fmt.Printf("%T\n\n", d3)

	fmt.Printf("RUNE\n")
	fmt.Printf("%T\n", e1)
	fmt.Printf("%T\n", e2)
	fmt.Printf("%T\n\n", e3)

	fmt.Printf("BYTE\n")
	fmt.Printf("%T\n", f1)
	fmt.Printf("%T\n", f2)
	fmt.Printf("%T\n\n", f3)
}
