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

// Comprueba que el tamaño inicial de la lista sea 0.
func TestInitialSizeMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	if arrayList.Size() != 0 {
		t.Fatalf("El tamaño inicial de arrayList debe ser 0, pero es %d", arrayList.Size())
	} else {
		fmt.Println("TestInitialSizeMethod: Tamaño inicial es 0")
	}
}

// Verifica que el método Add agregue elementos correctamente y que el tamaño de la lista sea el esperado.
func TestAddMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1
	e3 := 2

	arrayList.Add(e1)
	arrayList.Add(e2)
	arrayList.Add(e3)

	if arrayList.Size() != 3 {
		t.Fatalf("El método Add no funciona correctamente, el tamaño debe ser 3 y es %d", arrayList.Size())
	} else {
		fmt.Println("TestAddMethod: Los elementos se agregaron correctamente")
	}
}

// Verifica que el método AddAt agregue elementos en los índices correctos.
func TestAddAtMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1
	e3 := 2

	arrayList.Add(e1)
	arrayList.Add(e2)

	// Intentar agregar e3 en el índice 1
	if err := arrayList.AddAt(1, e3); err != nil {
		t.Fatalf("El método AddAt devolvió un error: %s", err)
	}

	if value, err := arrayList.Get(1); err != nil || value != e3 {
		t.Fatalf("El método AddAt no funciona correctamente, el valor del índice 1 debe ser %d", e3)
	}
}

// Comprueba que el método Get obtenga los elementos correctos en los índices especificados.
func TestGetMethod(t *testing.T) {
	arrayList := NewArrayList[int]()

	e1 := 0
	e2 := 1

	arrayList.Add(e1)
	arrayList.Add(e2)

	if value, err := arrayList.Get(0); err != nil || value != e1 {
		t.Fatalf("El método Get no funciona correctamente, get(0) debe ser 0, pero es %d", value)
	} else {
		fmt.Println("TestGetMethod: El valor en el índice 0 es correcto")
	}

	if value, err := arrayList.Get(1); err != nil || value != e2 {
		t.Fatalf("El método Get no funciona correctamente en iteración, get(1) debe ser 1, pero es %d", value)
	} else {
		fmt.Println("TestGetMethod: El valor en el índice 1 es correcto")
	}
}
