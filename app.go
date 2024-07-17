package main

import (
	ix "eda/structures/invertIndex"
	s "eda/types"

	// import trietree "eda/structures/trie"
	// import bplustree "eda/structures/bplus"

	"context"
	"fmt"
)

// App struct
type App struct {
	InvertIndex *ix.InvertIndex[s.Song]
	// Trie *trietree.Trie[s.Song]
	// Bplus *bplustree.Bplus[s.Song]
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// a.InvertIndex = ix.NewInvertIndex[s.Song](2000)
	// // a.Trie = trietree.NewTrie()
	// // a.Bplus = bplustree.NewBplus()

	// f := func(s *s.Song) {
	// 	name := s.TrackName
	// 	a.InvertIndex.PutMany(name, s)
	// 	// a.Trie.Add(name, s)
	// 	// a.Bplus.Add(name, s)
	// }

	// err := rd.ReadCSV(f, 1000)
	// if err != nil {
	// 	panic(err)
	// }
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
