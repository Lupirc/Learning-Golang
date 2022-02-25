package main

import "fmt"

func main() {
	//Map 1
	var myMap map[string]string
	myMap = map[string]string{}

	myMap["s1"] = "Lupi"
	myMap["s2"] = "Aji"

	fmt.Println(myMap)
	//fmt.Println(myMap["s1"])
	fmt.Println("=====================")

	//Map 2
	siswa := map[string]int{
		"Lupi":  9,
		"Brian": 7,
		"Sodik": 8,
	}

	for key, value := range siswa {
		fmt.Println("Key : ", key, " Value : ", value)
	}

	fmt.Println("=====================")

	//Menghapus Map
	delete(siswa, "Sodik")

	for _, value := range siswa {
		fmt.Println("Value : ", value)
	}

	fmt.Println("=====================")

	//Mengecek suatu key dalam map
	// cek := siswa["Sodik"]
	cek, aval := siswa["Lupi"]

	//fmt.Println(aval)
	//fmt.Println(cek)

	if aval {
		fmt.Println(cek)
	} else {
		fmt.Println(aval)
	}
}
