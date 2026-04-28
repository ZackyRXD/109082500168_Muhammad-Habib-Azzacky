package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(tab *tabel, n *int) {
	var karakter rune
	*n = 0
	fmt.Scanf("%c", &karakter)
	for karakter != '.' && *n < NMAX {
		tab[*n] = karakter
		*n = *n + 1
		fmt.Scanf("%c", &karakter)
	}
}

func cetakArray(tab tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(tab[i]))
	}
	fmt.Println()
}

func balikanArray(tab *tabel, n int) {
	var bantuan tabel
	for i := 0; i < n; i++ {
		bantuan[i] = tab[n-1-i]
	}
	for i := 0; i < n; i++ {
		tab[i] = bantuan[i]
	}
}

func palindrom(tab tabel, n int) bool {
	var asli tabel = tab
	balikanArray(&tab, n)
	for i := 0; i < n; i++ {
		if asli[i] != tab[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	balikanArray(&tab, n)
	cetakArray(tab, n)

	balikanArray(&tab, n)
	fmt.Println(palindrom(tab, n))
}