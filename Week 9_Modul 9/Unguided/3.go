package main

import "fmt"

func main() {
	var klub1 string
	var klub2 string
	var skor1 int
	var skor2 int
	var pemenang []string

	fmt.Print("Klub A : ")
	fmt.Scan(&klub1)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub2)

	var i int = 1
	for {
		fmt.Print("Pertandingan ", i, " : ")
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		if skor1 > skor2 {
			pemenang = append(pemenang, klub1)
		} else if skor2 > skor1 {
			pemenang = append(pemenang, klub2)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		i++
	}

	for j := 0; j < len(pemenang); j++ {
		fmt.Println("Hasil", j+1, ":", pemenang[j])
	}

	fmt.Println("Pertandingan selesai")
}