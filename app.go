package main

import (
	ix "eda/structures/invertIndex"
	s "eda/types"
	"time"

	t "eda/structures/trie"
	// import bplustree "eda/structures/bplus"
	"context"
	rd "eda/reader"
)

type App struct {
	InvertIndex *ix.InvertIndex[s.Song]
	Trie        *t.Trie[s.Song]
	// Bplus *bplustree.Bplus[s.Song]
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	a.InvertIndex = ix.NewInvertIndex[s.Song](2000)
	a.Trie = t.NewTrie[s.Song]()
	// a.Bplus = bplustree.NewBplus()

	f := func(s *s.Song) {
		name := s.TrackName
		a.InvertIndex.PutMany(name, s)
		a.Trie.Add(name, s)
		// a.Bplus.Add(name, s)
	}

	err := rd.ReadCSV(f, -1)
	if err != nil {
		panic(err)
	}
}

func (a *App) SearchSongInIndexInvert(name string) s.Result {
	start := time.Now()
	lists := a.InvertIndex.Search(name)

	timeLapse := float32(time.Since(start).Microseconds()) / 1000

	size := len(lists)

	if len(lists) > 500 {
		lists = lists[:500]
	}

	return s.Result{
		TimeLapse: timeLapse,
		Songs:     lists,
		Size:      size,
	}
}

func (a *App) SearchSongInTrie(name string) s.Result {
	start := time.Now()

	lists := a.Trie.Suggest(name).Parse()

	timeLapse := float32(time.Since(start).Microseconds()) / 1000

	size := len(lists)

	if len(lists) > 500 {
		lists = lists[:500]
	}

	return s.Result{
		TimeLapse: timeLapse,
		Songs:     lists,
		Size:      size,
	}
}
