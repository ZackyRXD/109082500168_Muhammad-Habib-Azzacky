package main
import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 0.8
}

func CelciusToFahrenheit(c suhu) suhu {
	return (c * 9 / 5) + 32
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var input suhu
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)
	fmt.Print("\n")

	fmt.Println(input, "celcius =", CelciusToReamur(input), "reamur")
	fmt.Println(input, "celcius =", CelciusToFahrenheit(input), "fahrenheit")
	fmt.Println(input, "celcius =", CelciusToKelvin(input), "kelvin")
}
