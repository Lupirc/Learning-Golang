package main

import "fmt"

func main() {
	// Percabangan IF
	nilai := 70
	var biji string

	if nilai >= 80 {
		biji = "A"
	} else if nilai < 80 && nilai >= 60 {
		biji = "B"
	} else if nilai < 60 && nilai >= 40 {
		biji = "C"
	} else {
		biji = "E"
	}

	fmt.Println(biji)

	//Percabangan Switch case
	menu := 2

	switch menu {
	case 1:
		fmt.Println("Menu 1")
	case 2:
		fmt.Println("Menu 2")
	case 3:
		fmt.Println("Menu 3")
	default:
		fmt.Println("Maaf Menu Tidak Ada")
	}

}
