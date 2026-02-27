package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
	/*
		start
		end
		3
		2
		1
	*/
}

// start печатается сразу
//end так же печатается сразу
//3 (defer'ы откладываются, и начинают выводить значения в порядке очереди LIFO)
//2
//1
