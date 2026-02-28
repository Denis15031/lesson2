package main

import "fmt"

func SaveDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Паника перехвачена: %v\n", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("деление на ноль")
	}
	return a / b
}

func main() {
	fmt.Println("SaveDivide(10, 2) =", SaveDivide(10, 2))
	fmt.Println("SaveDivide(10, 0) =", SaveDivide(10, 0))
}
