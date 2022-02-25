package main

import "fmt"

func main() {
	jumlah := jumlah(4)
	fmt.Println(jumlah)

	kalimat := hebat("Aku")
	fmt.Println(kalimat)

	tambah := tambah(20, 30)
	fmt.Println(tambah)
}

func jumlah(angka1 int) int {
	jml := angka1 + 5
	return jml
}

func hebat(sentence string) string {
	hebat := sentence + " HEBAT !!!"
	return hebat
}

func tambah(angka1 int, angka2 int) int {
	return angka1 + angka2
}
