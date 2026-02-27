package main

import (
	"errors"
	"fmt"
)

func SimpleError() error {
	return errors.New("простая ошибка")
}

func FormattedError(age int) error {
	return fmt.Errorf("ошибка: возраст %d недопустим: %w", age, errors.New("invalid age"))
}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("ошибка %d: %s", e.Code, e.Msg)
}

func StructError() error {
	return &MyError{Code: 404, Msg: "не найдено"}
}

func main() {
	fmt.Println("SimpleError:", SimpleError())

	fmt.Println("FormattedError(25):", FormattedError(25))

	err := StructError()
	fmt.Println("StructError:", err)

	if myErr, ok := err.(*MyError); ok {
		fmt.Printf("Код: %d, Сообщение: %s\n", myErr.Code, myErr.Msg)

	}
}
