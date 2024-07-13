package reader

import (
	invertI "eda/structures/invertIndex"
	s "eda/types"
	"fmt"
	"testing"
	"time"
)

func TestReadCSV(t *testing.T) {
	invertIndex := invertI.NewInvertIndex[s.Song](300 * 1000)

	f := func(s *s.Song) {
		name := s.TrackName

		invertIndex.PutMany(name, s)
	}

	err := ReadCSV(f, -1)
	if err != nil {
		t.Errorf("ReadCSV() failed : %v", err)
	}

	fmt.Println("Numero de claves: ", invertIndex.SizeKey())

	start := time.Now()

	list := invertIndex.Search("The People I Know")

	if list == nil {
		fmt.Println("Palara no encontrada")
		return
	}
	elapsed := time.Since(start)

	// list.ForEach(func(song *s.Song, i int) {
	// 	fmt.Println(i, song.TrackName, "id", song.TrackId)
	// })

	fmt.Println("len", list.Size())

	fmt.Println("Elapsed", elapsed.Microseconds(), "Microseconds")  // 0 microsegundos
	fmt.Println("Numero de colisiones", invertIndex.GetColliding()) // 38 011 colisiones
}
