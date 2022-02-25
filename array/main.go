package main

import "fmt"

func main() {
	//Aray 1
	var bahasa [5]string

	bahasa[0] = "Indo"
	bahasa[1] = "Jawa"
	bahasa[2] = "Sunda"
	bahasa[3] = "Kalimantan"
	bahasa[4] = "Madura"

	fmt.Println(bahasa)
	length1 := len(bahasa)
	fmt.Println(length1)

	//Array 2
	daerah := [5]string{"Indo", "Jawa", "Sunda", "Kalimantan", "Madura"}

	fmt.Println(daerah)
	length2 := len(daerah)
	fmt.Println(length2)

	//Array 3
	minuman := [...]string{
		"Boba",
		"Es Teh",
		"Jeruk",
		"Kopi",
	}

	for index, mimik := range minuman {
		fmt.Println("Index ke : ", index, " Isi : ", mimik)
	}
	length3 := len(minuman)
	fmt.Println(length3)
}
