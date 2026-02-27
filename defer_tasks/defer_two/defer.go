package main

import "fmt"

func main() {
	value := 123

	//Оборачиваем в функцию, захватывает переменную value по ссылке
	//Значение читается в момент выполнения defer
	defer func() {
		fmt.Println(value)
	}()

	changeValue(&value)
}

func changeValue(value *int) {
	*value = 456
}
