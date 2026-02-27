package main

import "fmt"

// Реализует интерфейс error
type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	return e.message
}

// Возвращает error, внутри которого может быть *CustomError
func returnError(flag bool) error {
	if flag {
		return &CustomError{"Что-то пошло не так"}
	}

	var err *CustomError // interface != nil!
	return err
}

func main() {
	err1 := returnError(true)  // Возвращает &CustomError{...}
	err2 := returnError(false) // Возвращает (type=*CustomError, value=nil)

	// err1: (type=*CustomError, value=&CustomError{...}) -> != nil
	fmt.Println("err1 == nil:", err1 == nil) // false

	// err2: (type=*CustomError, value=nil) -> != nil  (type не nil!)
	fmt.Println("err2 == nil:", err2 == nil) // false (ловушка!)
}
