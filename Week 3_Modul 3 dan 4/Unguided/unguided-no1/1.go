package main

import "fmt"

func hitungFaktorial(n int) int {
    hasil := 1
    for i := 1; i <= n; i++ {
        hasil *= i
    }
    return hasil
}

func cetakHasil(n, r int) {

	fN := hitungFaktorial(n)
    fNR := hitungFaktorial(n - r)
    fR := hitungFaktorial(r)

    p := fN / fNR
    c := fN / (fR * fNR)

    fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", n, r, p)
    fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", n, r, c)
}

func main() {
    var a, b, c, d int

    fmt.Print("masukkan nilai a : ")
    fmt.Scan(&a)
    fmt.Print("masukkan nilai b : ")
    fmt.Scan(&b)
    fmt.Print("masukkan nilai c : ")
    fmt.Scan(&c)
    fmt.Print("masukkan nilai d : ")
    fmt.Scan(&d)

    fmt.Println("--- Hasil Perhitungan ---")
    cetakHasil(a, c)
    cetakHasil(b, d)
}