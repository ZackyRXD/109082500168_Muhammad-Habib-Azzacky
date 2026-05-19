package main

import "fmt"

const KAPASITAS = 1000000

var daftar [KAPASITAS]int

func muatData(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&daftar[i])
	}
}

func temukan(n, nilai int) int {
	awal := 0
	akhir := n - 1
	for awal <= akhir {
		pertengahan := (awal + akhir) / 2
		if daftar[pertengahan] == nilai {
			return pertengahan
		} else if daftar[pertengahan] < nilai {
			awal = pertengahan + 1
		} else {
			akhir = pertengahan - 1
		}
	}
	return -1
}

func main() {
	var n, nilai int
	fmt.Scan(&n, &nilai)
	muatData(n)
	posisi := temukan(n, nilai)
	if posisi == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(posisi)
	}
}