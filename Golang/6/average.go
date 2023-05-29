package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Masukan Angka pisahkan dengan koma(,): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numbersStr := strings.Split(strings.TrimSpace(input), ",")
	numbers := make([]int, len(numbersStr))

	for i, numStr := range numbersStr {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			fmt.Printf("Angka tidak valid: %s\n", numStr)
			return
		}
		if num < 0 {
			fmt.Println("Masukan bilangan positif")
			return
		}
		numbers[i] = num
	}

	average := calculate(numbers)
	fmt.Printf("Rata-rata = %.2f\n", average)
}

func calculate(numbers []int) float64 {
	sum := 0

	for _, num := range numbers {
		sum += num
	}
	average := float64(sum) / float64(len(numbers))
	return average
}
