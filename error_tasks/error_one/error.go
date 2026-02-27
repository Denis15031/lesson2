package main

import "fmt"

func handle() error {
	return fmt.Errorf("Произошла ошибка в handle()")
}

func main() {
	err := handle()
	if err != nil {
		fmt.Println("Ошибка получена", err)
	}
}
