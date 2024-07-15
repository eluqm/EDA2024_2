package main

import ix "eda/structures/invertIndex"
import s "eda/types"
// import trietree "eda/structures/trie"
// import bplustree "eda/structures/bplus"
import rd "eda/reader"

type App struct {
  InvertIndex *ix.InvertIndex[s.Song]
  // Trie *trietree.Trie[s.Song] 
  // Bplus *bplustree.Bplus[s.Song] 
}

func InitApp() *App {
  app := App {
    InvertIndex: ix.NewInvertIndex[s.Song](2000),
    // Trie: trietree.NewTrie(),
    //Bplus: bplus.NewBplus(),
  }

  f := func(s *s.Song) {
    name := s.TrackName
    app.InvertIndex.PutMany(name, s)
    // app.Trie.Add(name, s)
    // app.Bplus.Add(name, s)
  }

  // Recibe como parametros: una función y un entero
  err := rd.ReadCSV(f, 1000)
  if (err != nil) {
    panic(err)
  }

  return &app
}

func main() {
  InitApp()
}
