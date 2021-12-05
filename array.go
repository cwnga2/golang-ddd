package main

import "fmt"

func main() {
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// var s []int = primes[1:4]
	// fmt.Println(s)
	// fmt.Println(primes)

	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }
	// fmt.Println(names)

	// c := names[0:2]
	// b := names[1:3]
	// fmt.Println(a, b)

	// b[0] = "XXX"
	// fmt.Println(c, b)
	// fmt.Println(names)

	var vv []int
	println(len(vv), cap(vv))

	// 添加一个空切片
	vv = append(vv, 10)
	println(len(vv), cap(vv))
	fmt.Printf("%v\n", vv)
	println(cap(vv))

}
