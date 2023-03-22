package main

import "fmt"

// Programa que inverte uma determinada string

func rev(s string) string {

	fraseInv := ""

	for _, v := range s {
		fraseInv = string(v) + fraseInv
	}
	return fraseInv
}

func main() {
	frase := "Garrafa" // Frase que vai ser invertida
	fmt.Println(rev(frase))
}
