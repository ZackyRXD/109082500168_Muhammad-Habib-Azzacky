package main

import "fmt"

func tampilkanGanjil(i int, n int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}

	tampilkanGanjil(i+1, n)
}

func main() {
	var n int

	fmt.Scan(&n)
	tampilkanGanjil(1, n)
}