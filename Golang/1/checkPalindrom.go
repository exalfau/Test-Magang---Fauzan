package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var kata string

func main() {
	fmt.Print("Masukan Kata: ")
	//input semua kata dalam kalimat
	reader := bufio.NewReader(os.Stdin)
	kata, _ = reader.ReadString('\n')
	//ubah kalimat menjadi lowercase dan hapus semua spasi
	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))
	//balik kalimat
	reversed := reverse(kata)
	kata = strings.TrimSpace(kata)
	reversed = strings.TrimSpace(reversed)
	palindrom := check(kata, reversed)
	fmt.Println("Palindrom: ", palindrom)
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func check(original, reversed string) bool {

	return original == reversed
}
