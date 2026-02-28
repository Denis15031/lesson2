package main

import "fmt"

func CausePaniс() {
	panic("что-то пошло не так!")
}

func HandlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Паника перехвачена: %v\n", r)
		}
	}()
	CausePaniс()
}

func main() {
	fmt.Println("Тест:обработанная паника")
	HandlePanic()
	fmt.Println("Программа продолжает выполнение!")
}
