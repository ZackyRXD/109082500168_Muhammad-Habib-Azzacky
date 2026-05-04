package main

import "fmt"

func main() {
    var n int
    var berat [1000]float64

    fmt.Print("mau berapa kelinci? : ")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        // PERBAIKAN: Gunakan koma untuk memisahkan teks dan variabel
        // Gunakan i+1 agar muncul "kelinci ke 1" bukan "kelinci ke 0"
        fmt.Print("masukan berat kelinci ke ", i+1, " : ") 
        fmt.Scan(&berat[i])
    }

    // Pastikan pengecekan min/max hanya jalan jika n > 0
    if n > 0 {
        min := berat[0]
        max := berat[0]

        for i := 1; i < n; i++ {
            if berat[i] < min {
                min = berat[i]
            }
            if berat[i] > max {
                max = berat[i]
            }
        }
        fmt.Println("Berat terkecil:", min, "Berat terbesar:", max)
    } else {
        fmt.Println("Data kosong.")
    }
}