package trie

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	trie := NewTrie[int]()

	e1 := 1

	fmt.Println("Agregando 'holaed'")
	trie.Add("holaed", &e1)
	fmt.Println("----")

	fmt.Println("Agregando 'hola'")
	trie.Add("hola", &e1)
	fmt.Println("----")

	fmt.Println("Agregando 'hol'")
	trie.Add("hol", &e1)
}

func TestAdd(t *testing.T) {
	trie := NewTrie[int]()
	e1 := 1

	words := []string{"hola", "holanda", "holar", "holas", "holan", "holant",
		"holari", "holaris", "holarith", "holarito", "holat", "holatero",
		"holario", "holarita", "holaring", "holanita", "holay", "holazo",
		"holal", "holamo", "holador", "holadura", "holadurao", "holaitis",
		"holando", "holament", "holancia", "holancha", "holandro", "holanza"}

	for _, word := range words {
		fmt.Printf("Agregando '%s'\n", word)
		trie.Add(word, &e1)
	}

	if *trie.Get("hola") != e1 {
		t.Fatalf("El valor esperado para 'hola' es %d", e1)
	}
}

func TestGet(t *testing.T) {
	trie := NewTrie[int]()
	e1 := 1

	words := []string{"hola", "holanda", "holar", "holas", "holan", "holant",
		"holari", "holaris", "holarith", "holarito", "holat", "holatero",
		"holario", "holarita", "holaring", "holanita", "holay", "holazo",
		"holal", "holamo", "holador", "holadura", "holadurao", "holaitis",
		"holando", "holament", "holancia", "holancha", "holandro", "holanza"}

	for _, word := range words {
		fmt.Printf("Agregando '%s'\n", word)
		trie.Add(word, &e1)
	}

	for _, word := range words {
		fmt.Printf("Obteniendo '%s'\n", word)
		if *trie.Get(word) != e1 {
			t.Fatalf("El valor esperado para '%s' es %d", word, e1)
		}
	}
}

func TestSearch(t *testing.T) {
	trie := NewTrie[int]()
	e1 := 1

	words := []string{"hola", "holanda", "holar", "holas", "holan", "holant",
		"holari", "holaris", "holarith", "holarito", "holat", "holatero",
		"holario", "holarita", "holaring", "holanita", "holay", "holazo",
		"holal", "holamo", "holador", "holadura", "holadurao", "holaitis",
		"holando", "holament", "holancia", "holancha", "holandro", "holanza"}

	for _, word := range words {
		fmt.Printf("Agregando '%s'\n", word)
		trie.Add(word, &e1)
	}

	for _, word := range words {
		fmt.Printf("Buscando '%s'\n", word)
		if !trie.Search(word) {
			t.Fatalf("'%s' debería existir en el trie", word)
		}
	}

	fmt.Println("Buscando 'mundo'")
	if trie.Search("mundo") {
		t.Fatalf("'mundo' no debería existir en el trie")
	}
}

func TestRemove(t *testing.T) {
	trie := NewTrie[int]()
	e1 := 1

	words := []string{"hola", "holanda", "holar", "holas", "holan", "holant",
		"holari", "holaris", "holarith", "holarito", "holat", "holatero",
		"holario", "holarita", "holaring", "holanita", "holay", "holazo",
		"holal", "holamo", "holador", "holadura", "holadurao", "holaitis",
		"holando", "holament", "holancia", "holancha", "holandro", "holanza"}

	for _, word := range words {
		fmt.Printf("Agregando '%s'\n", word)
		trie.Add(word, &e1)
	}

	for _, word := range words {
		fmt.Printf("Eliminando '%s'\n", word)
		value := trie.Remove(word)
		if value == nil || *value != e1 {
			t.Fatalf("El valor eliminado esperado para '%s' es %d", word, e1)
		}

		fmt.Printf("Search '%s' después de eliminar\n", word)
		if trie.Search(word) {
			t.Fatalf("'%s' no debería existir en el trie después de eliminar", word)
		}
	}
}

func TestSuggest(t *testing.T) {
	trie := NewTrie[int]()
	//Todavia hay bug con la ñ
	words := []string{"hola", "holandae", "holari", "holaso", "holanu", "holantq",
		"holariw", "holarisr", "holaritht", "holaritoy", "holatu", "holateroi",
		"holario", "holaritap", "holarings", "holanitad", "holayf", "holazg",
		"holalh", "holamoj", "holadork", "holadural", "holadurao", "holaitis",
		"holandoz", "holamentx", "holanciac", "holanchav", "holandrob", "holanzan"}

	for idx, word := range words {
		val := idx + 1
		fmt.Printf("Add '%s'\n", word)
		trie.Add(word, &val)
	}

	fmt.Println("Sugerencias para 'hol'")
	suggestions := trie.Suggest("hol")
	if len(suggestions) != 30 {
		t.Fatalf("Porfavorcito estoy aburrido, tiene que dar 30 sugerencias para 'hol', pero boto %d", len(suggestions))
	}
	for i := 0; i < 30; i++ {
		fmt.Println(*(suggestions[i]).GetValue(), string((suggestions[i]).GetKey()+'a'))
	}
}
