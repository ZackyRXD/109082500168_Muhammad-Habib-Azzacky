package main
import "fmt"
func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)
	if (tahun%4 == 0 && tahun%100 != 0) {
		kabisat = true
	} else {
		kabisat = false
	}
	if kabisat {
		fmt.Printf("%d adalah tahun kabisat.\n", tahun)
	} else {
		fmt.Printf("%d bukan tahun kabisat.\n", tahun)
	}
}