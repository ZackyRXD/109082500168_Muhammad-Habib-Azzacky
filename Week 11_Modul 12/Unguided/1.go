package main

import (
	"fmt"
)

func main() {
	rekap := make(map[int]int)
	masuk := 0
	sah := 0

	for {
		var pilihan int
		fmt.Scan(&pilihan)
		if pilihan == 0 {
			break
		}
		masuk++
		if pilihan >= 1 && pilihan <= 20 {
			sah++
			rekap[pilihan]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", masuk)
	fmt.Printf("Suara sah: %d\n", sah)

	for nomor := 1; nomor <= 20; nomor++ {
		if rekap[nomor] > 0 {
			fmt.Printf("%d: %d\n", nomor, rekap[nomor])
		}
	}
}