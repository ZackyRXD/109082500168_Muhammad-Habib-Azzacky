package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	
	fmt.Println("Masukkan jumlah ikan dan kapasitas wadah:")
	fmt.Scan(&x, &y)

	ikan := make([]float64, x)
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jmlWadah := int(math.Ceil(float64(x) / float64(y)))
	wadah := make([]float64, jmlWadah)
	var totalSemua float64

	for i := 0; i < x; i++ {
		indeks := i / y
		wadah[indeks] += ikan[i]
	}

	fmt.Println("Total berat per wadah:")
	for i := 0; i < jmlWadah; i++ {
		beratBulat := math.Round(wadah[i]*100) / 100
		fmt.Print(beratBulat)
		fmt.Print(" ")
		totalSemua += wadah[i]
	}
	fmt.Println()

	rataRata := totalSemua / float64(jmlWadah)
	rataRataBulat := math.Round(rataRata*100) / 100
	
	fmt.Println("Rata-rata berat seluruh wadah:")
	fmt.Println(rataRataBulat)
}