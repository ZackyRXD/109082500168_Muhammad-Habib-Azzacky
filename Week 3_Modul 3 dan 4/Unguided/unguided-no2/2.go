package main

import "fmt"

func kuadrat(x int) int {
	return x * x
}

func kurang(x int) int {
	return x - 2
}

func tambah(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai y : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai z : ")
	fmt.Scan(&c)

	
	hasilA := kuadrat(kurang(tambah(a)))
	fmt.Printf("f(g(h( %d ))) : %d\n", a, hasilA)

	hasilB := kurang(tambah(kuadrat(b)))
	fmt.Printf("g(h(f( %d ))) : %d\n", b, hasilB)

	hasilC := tambah(kuadrat(kurang(c)))
	fmt.Printf("h(f(g( %d ))) : %d\n", c, hasilC)
}