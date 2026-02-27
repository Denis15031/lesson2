package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

var (
	ErrNotFound  = errors.New("ресурс не найден")
	TimeoutError = errors.New("таймаут операции")
)

// имитация запроса с разными ошибками
func SimulateRequest() error {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return errors.New("ошибка генерации случайного числа")
	}

	r := int(n.Int64())

	switch {
	case r < 50:
		return fmt.Errorf("запрос не выполнен: %w", TimeoutError)
	case r < 80:
		return fmt.Errorf("ошибка: %w", ErrNotFound)
	default:
		return errors.New("неизвестная ошибка")
	}
}

//анализируем цепочку ошибок

func ProcessError(err error) {
	if errors.Is(err, TimeoutError) {
		fmt.Println("Требуется повторная попытка")
		return
	}

	if errors.Is(err, ErrNotFound) {
		fmt.Println("Ресурс не найден")
		return
	}

	fmt.Println("Неизвестная ошибка")
}

func main() {
	fmt.Println("Тестирование ProcessError")
	for i := 0; i < 5; i++ {
		err := SimulateRequest()
		fmt.Printf("Запрос #%d: %v ->", i+1, err)
		ProcessError(err)
	}
}
