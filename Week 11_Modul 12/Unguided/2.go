package main

import (
	"fmt"
)

func main() {
	perolehan := make(map[int]int)
	totalMasuk := 0
	totalSah := 0

	for {
		var suara int
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		totalMasuk++
		if suara >= 1 && suara <= 20 {
			totalSah++
			perolehan[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalMasuk)
	fmt.Printf("Suara sah: %d\n", totalSah)

	calonKetua := -1
	nilaiKetua := -1
	for no := 1; no <= 20; no++ {
		if perolehan[no] > nilaiKetua {
			nilaiKetua = perolehan[no]
			calonKetua = no
		}
	}

	calonWakil := -1
	nilaiWakil := -1
	for no := 1; no <= 20; no++ {
		if no == calonKetua {
			continue
		}
		if perolehan[no] > nilaiWakil {
			nilaiWakil = perolehan[no]
			calonWakil = no
		}
	}

	fmt.Printf("Ketua RT: %d\n", calonKetua)
	fmt.Printf("Wakil ketua: %d\n", calonWakil)
}