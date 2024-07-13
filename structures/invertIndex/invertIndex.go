package invertindex

import (
	l "eda/structures/list/linkedList"
	h "eda/structures/map/hashMap"
	"fmt"
	"strings"
	"sync"
	"time"
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
	key = strings.ToLower(key)

	list := index.hashMap.Get(key)

	if list == nil {
		list = l.NewLinkedList[T](func(a, b T) bool {
			return &a == &b
		})
		index.hashMap.Put(key, list)
	}

	list.AddFirst(value)
	index.size = index.size + 1
}

func (index *InvertIndex[T]) Remove(key string) {
	index.hashMap.Remove(key)
	index.size = index.size - 1
}

func (index *InvertIndex[T]) Size() int {
	return index.size
}

func (index *InvertIndex[T]) SizeKey() int {
	return index.hashMap.Size()
}

func (index *InvertIndex[T]) IsEmpty() bool {
	return index.size == 0
}

func (index *InvertIndex[T]) PutMany(str string, value *T) {
	words := strings.Split(strings.ToLower(str), " ")

	for _, word := range words {
		word = ClearWord(word)
		index.Put(word, value)
	}
}

func ClearWord(word string) string {
	var sb strings.Builder

	for i, l := range word {
		if l >= 'a' && l <= 'z' {
			sb.WriteRune(l)
		} else if l == 164 {
			sb.WriteRune('ñ')
		}

		if i > (len(word)*2)/4 {
			break
		}
	}
	return sb.String()
}

func (index *InvertIndex[T]) GetColliding() int {
	return index.hashMap.Colliding
}

func (index *InvertIndex[T]) Search(str string) *l.LinkedList[T] {
	words := strings.Split(str, " ")

	start := time.Now()

	for _, word := range words {
		l := index.Get(word)
		fmt.Println("array encontrado con ", l.Size(), word)
	}
	elapsed := time.Since(start)

	fmt.Println("Demoro en buscar ", len(words), " palabras ", elapsed.Microseconds(), "Microseconds")

	return nil
}

func (index *InvertIndex[T]) Get(key string) *l.LinkedList[T] {
	key = strings.ToLower(key)

	key = ClearWord(key)

	return index.hashMap.Get(key)
}

func (index *InvertIndex[T]) AsyncSearch(str string) *l.LinkedList[T] {
	words := strings.Split(str, " ")
	var wg sync.WaitGroup

	start := time.Now()
	results := make([]*l.LinkedList[T], len(words))

	for i, word := range words {
		wg.Add(1)
		go func(i int, word string) {
			defer wg.Done()
			results[i] = index.Get(word)
			fmt.Println("array encontrado con ", results[i].Size(), word)
		}(i, word)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("Demoro en buscar ", len(words), " palabras ", elapsed.Microseconds(), "Microseconds")

	// Aquí podrías combinar los resultados de alguna manera
	return nil
}
