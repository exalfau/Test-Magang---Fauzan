package main

import "fmt"

var angka int

func main() {
	fmt.Print("Masukan Angka: ")
	fmt.Scanln(&angka)
	output := factor(angka)
	if angka < 0 {
		fmt.Println("Masukan bilangan positif")
	} else {
		fmt.Printf("%d! = %d ", angka, output)
	}
}

func factor(angka int) int {
	if angka < 1 {
		return 1
	}

	return angka * factor(angka-1)
}
