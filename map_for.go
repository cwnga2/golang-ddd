package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    strArray := strings.Fields(s)
	var result = map[string] int{}
    for _, stringValue := range strArray {
        result[stringValue]++
    }
    return result
}

func main() {
    wc.Test(WordCount)
}

