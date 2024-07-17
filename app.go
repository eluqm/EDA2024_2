package main

import (
	ix "eda/structures/invertIndex"
	s "eda/types"

	// import trietree "eda/structures/trie"
	// import bplustree "eda/structures/bplus"
	"context"
	rd "eda/reader"
	"fmt"
)

type App struct {
	InvertIndex *ix.InvertIndex[s.Song]
	// Trie *trietree.Trie[s.Song]
	// Bplus *bplustree.Bplus[s.Song]
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	a.InvertIndex = ix.NewInvertIndex[s.Song](2000)
	// a.Trie = trietree.NewTrie()
	// a.Bplus = bplustree.NewBplus()

	f := func(s *s.Song) {
		name := s.TrackName
		a.InvertIndex.PutMany(name, s)
		// a.Trie.Add(name, s)
		// a.Bplus.Add(name, s)
	}

	err := rd.ReadCSV(f, -1)
	if err != nil {
		panic(err)
	}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SearchSong(name string) []s.Song {
	lists := a.InvertIndex.Search(name)

	if len(lists) > 1000 {
		return lists[:1000]
	}

	return lists
}
