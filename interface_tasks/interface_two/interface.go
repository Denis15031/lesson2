package main

import "fmt"

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

// сравниваем err с nil
func checkErr(err error) {
	fmt.Println(err == nil)
}

func main() {
	var e1 error
	checkErr(e1) //true// -zero value для interface = (nil, nil)

	var e *errorString
	checkErr(e) //false// nil указатель, type != nil

	e = &errorString{}
	checkErr(e) //false// не nil указатель, type, value != nil

	e = nil
	checkErr(e) //false // присваиваем nil, но тип *errorString
}
