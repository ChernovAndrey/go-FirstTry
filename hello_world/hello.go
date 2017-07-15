package main

import (
	"fmt"
)

func makeEvenGenerator() func() (uint, uint) {
	var i uint = 0
	fmt.Println("i=", i)
	return func() (ret, i1 uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4
}
