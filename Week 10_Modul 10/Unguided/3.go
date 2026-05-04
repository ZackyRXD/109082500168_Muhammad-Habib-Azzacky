package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, min, max *float64) {
	*min = arrBerat[0]
	*max = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *min {
			*min = arrBerat[i]
		}
		if arrBerat[i] > *max {
			*max = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var data arrBalita
	var beratMin, beratMax float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&data[i])
	}

	hitungMinMax(data, n, &beratMin, &beratMax)
	rata := rerata(data, n)

	fmt.Println("Berat balita minimum:", beratMin, "kg")
	fmt.Println("Berat balita maksimum:", beratMax, "kg")
	fmt.Println("Rerata berat balita:", rata, "kg")
}