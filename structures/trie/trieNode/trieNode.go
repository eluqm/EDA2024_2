package trieNode

import "fmt"

type TrieNode[T any] struct {
	key    byte
	value  *T
	childs []*TrieNode[T]
	end    bool
}

func NewTrieNode[T any](end bool, value *T, key byte) *TrieNode[T] {
	return &TrieNode[T]{
		childs: make([]*TrieNode[T], 27),
		end:    end,
		value:  value,
		key:    key,
	}
}

func (node *TrieNode[T]) GetValue() *T {
	return node.value
}

func (node *TrieNode[T]) GetInNode(bytes *[]byte, i int) *T {
	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 && node.childs[key].end {
		return node.childs[key].value
	}

	return node.childs[key].GetInNode(bytes, i+1)
}

func (node *TrieNode[T]) GetKey() byte {
	return node.key
}

func (node *TrieNode[T]) AddInNode(bytes *[]byte, i int, value *T) {
	key := (*bytes)[i]
	isFinal := len(*bytes) == i+1

	if node.childs[key] == nil {
		fmt.Println(" - No existe")
		if isFinal {
			node.childs[key] = NewTrieNode[T](true, value, key)
		} else {
			fmt.Println(" - No es final")
			node.childs[key] = NewTrieNode[T](false, nil, key)
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	} else {
		fmt.Println(" - Existe")
		if isFinal {
			fmt.Println(" - Es final y se repite")
			node.childs[key].end = true
			node.childs[key].value = value
		} else {
			fmt.Println(" - No es final")
			node.childs[key].AddInNode(bytes, i+1, value)
		}
	}
}

func (node *TrieNode[T]) SearchInNode(bytes *[]byte, i int) bool {
	// Chequear si el índice está dentro del rango del array de bytes
	if i >= len(*bytes) {
		return false
	}

	key := (*bytes)[i]

	if node.childs[key] == nil {
		fmt.Println("No existe")
		return false
	}

	if len(*bytes) == i+1 && node.childs[key].end {
		return true
	}

	return node.childs[key].SearchInNode(bytes, i+1)
}

func (node *TrieNode[T]) RemoveInNode(bytes *[]byte, i int) *T {
	// Chequear si el índice está dentro del rango del array de bytes
	if i >= len(*bytes) {
		return nil
	}

	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 && node.childs[key].end {
		node.childs[key].end = false

		value := node.childs[key].value
		node.childs[key].value = nil

		return value
	}

	return node.childs[key].RemoveInNode(bytes, i+1)
}

func (node *TrieNode[T]) SearchPreFix(bytes *[]byte, i int) *TrieNode[T] {
	key := (*bytes)[i]

	if node.childs[key] == nil {
		return nil
	}

	if len(*bytes) == i+1 {
		return node.childs[key]
	}

	return node.childs[key].SearchPreFix(bytes, i+1)
}

func (node *TrieNode[T]) GetAllChild(suggest *[]*TrieNode[T]) {
	if node.end {
		*suggest = append(*suggest, node)
	}

	for _, child := range node.childs {
		if child != nil {
			child.GetAllChild(suggest)
		}
	}
}
