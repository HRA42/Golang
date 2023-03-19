package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	t, z := 0., 1.
	for i := 1; ; i++ {
		z, t = z-(z*z-x)/(2*z), z
		fmt.Printf("Die %d. Schätzung ist: %e \n", i, z)
		if math.Abs(z-t) < 1e-2 {
			break
		}
	}
	return z
}

func main() {
	x := 42.
	fmt.Printf("Das Ergebnis mit der selbst erstellten Funktion ist: %e \n", Sqrt(x))
	fmt.Printf("Ergebnis mit math.sqrt: %e \n", math.Sqrt(x))
}
