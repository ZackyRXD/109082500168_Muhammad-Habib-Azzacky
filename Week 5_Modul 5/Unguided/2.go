package main

import "fmt"

func rekursif(i, n int) {
	if i > n {
		return
	}

	if i%2 != 0 {
		fmt.Printf("%d ", i)
	}

	rekursif(i+1, n)
}

func main() {
	var x int
	fmt.Scan(&x)
	rekursif(1, x)
	fmt.Println()
}