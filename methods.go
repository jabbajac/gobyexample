package main

import "fmt"

type rect struct {
	length, height int
}

func (r *rect) area() int {
	return r.length * r.height
}

func (r *rect) perimeter() int {
	return 2*r.length + 2*r.height
}

func main() {
	r := rect{length: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perimeter())
}
