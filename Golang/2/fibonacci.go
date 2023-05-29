package main

import (
	"fmt"
)

var angka int

func main() {
	fmt.Print("Masukan Angka: ")
	fmt.Scanln(&angka)
	output := fibo(angka)
	fmt.Println(output)

}

func fibo(angka int) []int {
	fibonacci := make([]int, angka)
	fibonacci[0] = 0
	fibonacci[1] = 1

	for i := 2; i < angka; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci
}
