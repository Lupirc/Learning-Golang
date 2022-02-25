package main

import (
	"fmt"
	"test/kalkulator"
)

func main() {
	fmt.Println("Aplikasi Kalkulator")

	// Definisi variabel cara 1
	var resultTambah int
	var resultKurang int

	resultTambah = kalkulator.Add(1, 4)
	resultKurang = kalkulator.Sub(5, 2)

	//Definisi Variabel cara 2 dengan langsung mengisi variabel
	resultKali := kalkulator.Mul(4, 2)
	resultBagi := kalkulator.Div(10, 2)

	fmt.Println(resultTambah)
	fmt.Println(resultKurang)
	fmt.Println(resultKali)
	fmt.Println(resultBagi)
}
