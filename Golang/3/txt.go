package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Print("Masukan Nama File: ")
	var filename string
	fmt.Scanln(&filename)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Gagal Membaca File : %v", err)
	}
	fmt.Println(string(content))
}
