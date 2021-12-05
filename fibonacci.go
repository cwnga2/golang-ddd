package main

import "fmt"

// 返回一个“返回int的函数”
func tmp(step int) int {
    if step == 0 || step == 1 {
        return step
    } else {
        return tmp(step - 1) + tmp(step-2)
    }
}
func fibonacci() func() int {

    var i = 0
    return func() int {

        var result = tmp(i)
        i++
        return result
    }
}

func main() {
    f := fibonacci()
    for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}

