package main

import "fmt"

// самый глубокий уровень, вызывает панику
func Level3() {
	panic("ошибка в Level3")
}

// средний уровень, регистрирует defer для логирования
func Level2() {
	defer func() {
		fmt.Println("Завершаем Level2")
	}()

	Level3()
}

// верхний уровень, перехватывает панику через рековер
func Level1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Паника отработана на уровне 1: %v\n", r)
		}
	}()

	Level2()
}

func main() {
	fmt.Println("Многоуровневая паника")
	Level1()
	fmt.Println("Программа завершена корректно")
}
