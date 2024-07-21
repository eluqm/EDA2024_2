package arraylist

import (
// "errors"
// "fmt"
)

type ArrayList[T any] struct {
	elements []T
}

// Crea y devuelve una nueva instancia de ArrayList.
func NewArrayList[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		elements: make([]T, 0),
	}
}
