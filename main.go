package main

import "fmt"

func main() {
	var limite_ciclo float64
	var n float64 = 1.0
	var resultado float64 = 0
	var e float64
	fmt.Scanln(&limite_ciclo)
	for n <= limite_ciclo {
		e = 1.0 / n
		resultado = resultado + e
		n++
	}
	fmt.Println(resultado)
}
