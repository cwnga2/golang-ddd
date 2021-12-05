package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	println(a.X)
	println(a.Y)

	v := Vertex{1, 2}
	p := &v
	pp := v
	v.Y = 999999
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(pp)

}
