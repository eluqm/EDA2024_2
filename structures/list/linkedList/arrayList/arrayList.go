package arraylist

import (
	"errors"
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

// Devuelve el número de elementos en la lista.
func (list *ArrayList[T]) Size() int {
	return len(list.elements)
}

// Agrega un nuevo elemento a la lista.
func (list *ArrayList[T]) Add(data T) {
	list.elements = append(list.elements, data)
}

// Devuelve el elemento en el índice especificado.
func (list *ArrayList[T]) Get(index int) (T, error) {
	var zeroValue T
	if index < 0 || index >= len(list.elements) {
		return zeroValue, errors.New("index out of limits")
	}
	return list.elements[index], nil
}

// Agrega un nuevo elemento en un índice específico.
func (list *ArrayList[T]) AddAt(index int, data T) error {
	if index < 0 || index > len(list.elements) {
		return errors.New("index out of limits")
	}
	list.elements = append(list.elements[:index], append([]T{data}, list.elements[index:]...)...)
	return nil
}

// Establece el elemento en el índice especificado.
func (list *ArrayList[T]) Set(index int, data T) error {
	if index < 0 || index >= len(list.elements) {
		return errors.New("index out of limits")
	}
	list.elements[index] = data
	return nil
}

// Elimina el elemento en el índice especificado.
func (list *ArrayList[T]) Remove(index int) (T, error) {
	var zeroValue T
	if index < 0 || index >= len(list.elements) {
		return zeroValue, errors.New("index out of limits")
	}
	removed := list.elements[index]
	list.elements = append(list.elements[:index], list.elements[index+1:]...)
	return removed, nil
}
