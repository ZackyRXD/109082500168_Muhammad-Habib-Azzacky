package main
import "fmt"
func main() {
	var totalGram, sisaGram, kg, biayaKG, biayaSisa int
	fmt.Print ("Masukan berat parsel (gram): ")
	fmt.Scan(&totalGram)

	kg = totalGram / 1000
	sisaGram = totalGram % 1000
	biayaKG = kg * 10000
	
	if kg >= 10 { biayaSisa = 0 } else if sisaGram >= 500 {
		biayaSisa = sisaGram * 5 
	} else if sisaGram < 500 {
		biayaSisa = sisaGram * 15
	}

	fmt.Println("Detail berat: " , kg , " kg dan ", sisaGram, " gram")
	fmt.Println("Detail biaya: " ,"RP" , biayaKG , " + " , "RP" , biayaSisa)
	fmt.Println("Total biaya : " , "RP" , biayaKG + biayaSisa) 

	
}