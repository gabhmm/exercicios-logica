package main

import (
	"fmt"
	"time"
)

// programa que mostra os meses do ano e os dias de cada um deles usando sistema de loop

func main() {
	for m := 1; m <= 12; m++ {
		fmt.Println("Mês:", time.Month(m))

		if m <= 7 && m%2 != 0 {
			for d := 1; d <= 31; d++ {
				fmt.Print("Dia: ", d, "\t")
			}
		} else if m == 2 {
			for d := 1; d <= 28; d++ {
				fmt.Print("Dia: ", d, "\t")
			}
		} else if m <= 7 && m%2 == 0 {
			for d := 1; d <= 30; d++ {
				fmt.Print("Dia: ", d, "\t")
			}
		} else if m >= 8 && m%2 != 0 {
			for d := 1; d <= 30; d++ {
				fmt.Print("Dia: ", d, "\t")
			}
		} else if m >= 8 && m%2 == 0 {
			for d := 1; d <= 31; d++ {
				fmt.Print("Dia: ", d, "\t")
			}
		} else {
			fmt.Print("Nill")
		}
		fmt.Println("")
	}

}
