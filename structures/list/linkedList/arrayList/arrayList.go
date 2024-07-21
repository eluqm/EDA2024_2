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

// Verifica si la lista contiene el elemento especificado.
func (list *ArrayList[T]) Contains(data T, equals func(a, b T) bool) bool {
	for _, e := range list.elements {
		if equals(e, data) {
			return true
		}
	}
	return false
}

// Devuelve el índice del primer elemento que coincide con el especificado.
func (list *ArrayList[T]) IndexOf(data T, equals func(a, b T) bool) int {
	for i, e := range list.elements {
		if equals(e, data) {
			return i
		}
	}
	return -1
}

// IsEmpty verifica si la lista está vacía.
func (list *ArrayList[T]) IsEmpty() bool {
	return len(list.elements) == 0
}
