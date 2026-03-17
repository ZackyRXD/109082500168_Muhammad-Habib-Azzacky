package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	
	fmt.Println(1)
}

func main() {
	var bilanganAwal int

	fmt.Print("Masukkan bilangan : ")
	_, err := fmt.Scan(&bilanganAwal)
	
	if err != nil || bilanganAwal <= 0 || bilanganAwal >= 1000000 {
		return
	}

	cetakDeret(bilanganAwal)
}