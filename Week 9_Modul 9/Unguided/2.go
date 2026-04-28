package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var x int
	var hapus int
	var cari int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	var arr []int = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Isi elemen ke-", i, ": ")
		fmt.Scan(&arr[i])
	}

	fmt.Println("a. Isi array:", arr)

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("d. Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Elemen indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var sum float64 = 0
	for _, v := range arr {
		sum += float64(v)
	}
	var rata float64 = sum / float64(len(arr))
	fmt.Println("f. Rata-rata:", rata)

	var varians float64
	for _, v := range arr {
		varians += math.Pow(float64(v)-rata, 2)
	}
	var stdev float64 = math.Sqrt(varians / float64(len(arr)))
	fmt.Println("g. Standar Deviasi:", stdev)

	fmt.Print("h. Cari frekuensi angka: ")
	fmt.Scan(&cari)
	var count int = 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi", cari, ":", count, "kali")

	fmt.Print("e. Indeks yang ingin dihapus: ")
	fmt.Scan(&hapus)
	arr = append(arr[:hapus], arr[hapus+1:]...)
	fmt.Println("Isi array setelah dihapus:", arr)
}