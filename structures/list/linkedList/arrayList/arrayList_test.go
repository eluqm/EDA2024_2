package arraylist

import (
	"fmt"
	"testing"
)

// Verifica que el constructor NewArrayList inicialice bien la lista.
func TestNewArrayList(t *testing.T) {
	arrayList := NewArrayList[int]()

	if arrayList == nil {
		t.Fatal("ArrayList struct no inicializada")
	} else {
		fmt.Println("TestNewArrayList: ArrayList inicializada correctamente")
	}
}
