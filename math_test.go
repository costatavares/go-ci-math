package main

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {

	total := Soma(15, 15)

	esperado := 35

	if total != esperado {
		t.Errorf("Resultado da soma é invalido: Resultado %d.  Esperado %d.", total, esperado)
	} else {
		fmt.Println("Resultado: 30")
	}
}
