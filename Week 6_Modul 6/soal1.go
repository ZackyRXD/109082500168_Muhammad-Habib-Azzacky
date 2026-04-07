package main

import "fmt"

func volume(r, t float64) float64 {
	return 3.14 * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 > m2 {
		fmt.Println("Selisih massa zat cair kiri dan zat cair kanan :", m1-m2)
	} else {
		fmt.Println("Selisih massa zat cair kiri dan zat cair kanan :", m2-m1)
	}
}

func main() {
	var r, tKiri, tKanan, mjKiri, mjKanan float64

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	display(massa(r, tKiri, mjKiri), massa(r, tKanan, mjKanan))
}