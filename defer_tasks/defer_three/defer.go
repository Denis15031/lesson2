package main

import (
	"errors"
	"fmt"
)

func main() {
	println("Case 1") //Case 1
	case1()
	println()
	println()

	println("Case 2") //Case 2
	case2()
	println()
	println()

	println("Case 3") // Case 3
	case3()
	println()
	println()

}

func case1() {
	//функция с Defer, но локальная переменная retVal
	helperWithDefer := func(isError bool) error {
		var retVal error // error интерфейс, zero value = nil, retVal = nil по умолчанию

		defer func() {
			retVal = errors.New("Extra error") //выполнится после копирования return, здесь изменения не влияет на возвращаемое значение
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal //значение копируется здесь до выполнения defer
	}

	helperWithoutDefer := func(isError bool) error {
		var retVal error

		if isError {
			retVal = errors.New("Default error")
		}

		return retVal
	}

	fmt.Println("\twithout:")              //without:
	fmt.Println(helperWithoutDefer(false)) // nil // error интерфейс, zero value = nil
	fmt.Println(helperWithoutDefer(true))  // Default error //isError = true -> retVal= "Default error" -> return "Default error"
	fmt.Println("\twith:")                 // with:
	fmt.Println(helperWithDefer(false))    //nil //return скопировал nil до выполнения defer
	fmt.Println(helperWithDefer(true))     // Default error //return скопировал "Default" до выполнения defer
}

func case2() {
	//Функция с именованным return параметром
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() {
			retVal = errors.New("Extra error") // с именованным return defer выполняется перед фактическим возвратом
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return //defer выполнится после, но до фактического возвращения из функции
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return // вернет текущее значение retVal
	}

	fmt.Println("\twithout:")              // without:
	fmt.Println(helperWithoutDefer(false)) // nil // именованный параметр инициализируется zero value
	fmt.Println(helperWithoutDefer(true))  // Default error // логика та же как и в первом кейсе здесь
	fmt.Println("\twith:")                 // with:
	fmt.Println(helperWithDefer(false))    // Extra error //defer успевает изменить retVal до возврата
	fmt.Println(helperWithDefer(true))     // Extra error // defer перезаписывает Default на Extra
}

func case3() {
	//функция с двумя Defer
	helperWithDefer := func(isError bool) (retVal error) {
		defer func() { // выполнится вторым
			retVal = errors.New("First Error")
		}()

		defer func() { //выполнится первым
			retVal = errors.New("Second Error")
		}()

		if isError {
			retVal = errors.New("Default error")
		}

		return // вернет текущее значение retVal, defer выполнится после
	}

	helperWithoutDefer := func(isError bool) (retVal error) {
		if isError {
			retVal = errors.New("Default error")
		}

		return
	}

	fmt.Println("\twithout:")              // wuthout:
	fmt.Println(helperWithoutDefer(false)) //nil //  is.Error=false -> retVal=nil -> return nil
	fmt.Println(helperWithoutDefer(true))  //Default error // логика та же, как в 1 и 2 кейсе
	fmt.Println("\twith:")                 // with:
	fmt.Println(helperWithDefer(false))    // First Error //  defer 1 выполнится последним и перезапишет Second
	fmt.Println(helperWithDefer(true))     // First Error //  defer 1 - последнее изменение перед возвратом
}

/*
Case 1
	without:
<nil>
Default error
	with:
<nil>
Default error


Case 2
	without:
<nil>
Default error
	with:
Extra error
Extra error


Case 3
	without:
<nil>
Default error
	with:
First Error
First Error
*/
