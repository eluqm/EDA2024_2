package invertindex

import (
	l "eda/structures/list/linkedList"
	h "eda/structures/map/hashMap"
)

type InvertIndex[T any] struct {
	hashMap *h.HashMap[l.LinkedList[T]]

	size int
}

func NewInvertIndex[T any](capacity int) *InvertIndex[T] {
	const (
		SIZE int = 0
	)

	return &InvertIndex[T]{
		hashMap: h.NewHashMap[l.LinkedList[T]](capacity),
		size:    SIZE,
	}
}

func (index *InvertIndex[T]) Put(key string, value *T) {
	list := index.hashMap.Get(key)

	if list == nil {
		list = l.NewLinkedList[T](func(a, b T) bool {
			return &a == &b
		})
		index.hashMap.Put(key, list)
	}

	list.Add(value)
	index.size = index.size + 1
}

func (index *InvertIndex[T]) Get(key string) *l.LinkedList[T] {
	return index.hashMap.Get(key)
}

func (index *InvertIndex[T]) Remove(key string) {
	index.hashMap.Remove(key)
	index.size = index.size - 1
}

func (index *InvertIndex[T]) Size() int {
	return index.size
}

func (index *InvertIndex[T]) IsEmpty() bool {
	return index.size == 0
}
