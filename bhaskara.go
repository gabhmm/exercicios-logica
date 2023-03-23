package main

import (
	"fmt"
	"math"
)

//programa para calcular formula de bhaskara e mostrar valores na tela.

var a float64
var b float64
var c float64

func main() {

	a = 2
	b = 12
	c = -14

	delta := (math.Pow(b, 2) - (4 * a * c))

	if delta <= 0 {
		fmt.Printf("Essa equação não possui raízes reais, pois Δ = %v", delta)

	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("Δ = %v\n", delta)
		fmt.Printf("X1 = %v\nX2 = %v", x1, x2)
	}

}
