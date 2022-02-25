package main

import "fmt"

func main() {
	//Slice 1 rekomen
	var siswa []string

	siswa = append(siswa, "Aji")
	siswa = append(siswa, "Lupi")
	siswa = append(siswa, "Mung")

	//fmt.Println(siswa)
	for _, nama := range siswa {
		fmt.Println(nama)
	}
	length := len(siswa)
	fmt.Println(length)

	//Slice 2 Rekomen
	sekolah := []string{
		"Smea",
		"Smkk",
		"Cawang",
	}
	fmt.Println(sekolah)
}
