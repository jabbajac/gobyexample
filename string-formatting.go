package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// prints an instance of the point struct
	fmt.Printf("%v\n", p)
	// if the value is a struct, +v prints the struct's field names
	fmt.Printf("%+v\n", p)
	// prints go syntax representation of the value
	fmt.Printf("%#v\n", p)
	// prints the type of the value
	fmt.Printf("%T\n", p)
	// formatting booleans
	fmt.Printf("%t\n", true)
	// %d formats base 10 for integer
	fmt.Printf("%d\n", 123)
	// %b prints binary representation
	fmt.Printf("%b\n", 14)
	// prints the character representation of the integer
	fmt.Printf("%c\n", 33)
	// %x prints hex character
	fmt.Printf("%x\n", 456)
	// %f prints float
	fmt.Printf("%f\n", 78.9)
	// %e and %E format float in slightly different versions of scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// use %s for basic string printing
	fmt.Printf("%s\n", "\"string\"")
	// to double quote use %q
	fmt.Printf("%q\n", "\"string\"")
	// %x renders string in base 16 with 2 chars per byte
	fmt.Printf("%x\n", "hex this")
	// %p prints a pointer
	fmt.Printf("%p\n", &p)
	// you can use %#d to sepcify float width and precision
	fmt.Printf("|%6d|%5d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

}
