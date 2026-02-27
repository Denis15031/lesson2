package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0), //пустой слайс
	}
}

// Добавляет элемент на вершину стэка
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Удаляет и возвращает верхний элемент
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.elements) - 1

	value := s.elements[index]

	s.elements = s.elements[:index]

	return value, true
}

// Возвращает верхний элемент без удаления
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

// Проверяет, пуст ли стек
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	// Стек для int
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)

	if val, ok := intStack.Pop(); ok {
		fmt.Println("Pop int:", val) // 20
	}

	// Стек для string
	strStack := NewStack[string]()
	strStack.Push("hello")
	strStack.Push("world")

	if val, ok := strStack.Peek(); ok {
		fmt.Println("Peek string:", val) // "world"
	}

	// Попытка Pop из пустого стека
	emptyStack := NewStack[float64]()
	if _, ok := emptyStack.Pop(); !ok {
		fmt.Println("Попытка Pop из пустого стека — безопасно!")
	}
}
