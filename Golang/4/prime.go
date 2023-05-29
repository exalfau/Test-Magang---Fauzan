package main

import (
	"fmt"
)

var angka int

func main() {
	fmt.Print("Masukan Angka: ")
	fmt.Scanln(&angka)
	if angka > 0 {
		output := prime(angka)
		fmt.Println(output)
	} else {
		fmt.Println("Masukan Angka Positif")
	}
}

func prime(limit int) []int {
	var primes []int

	for num := 2; num < limit; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, num)
		}
	}

	return primes
}
