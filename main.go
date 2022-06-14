package main

import (
	"fmt"

	matematika "github.com/brnhrdwnnr/cek-ganjil-genap"
	matematikaV2 "github.com/brnhrdwnnr/cek-ganjil-genap/v2"
)

func main() {
	result := matematika.CekGanjilGenap(1)
	result2 := matematikaV2.CekGanjilGenap(1, 2, 3, 4, 5)
	fmt.Println("V1: " + result)
	fmt.Println("V2: " + result2)

}