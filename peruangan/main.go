package main

import (
	"fmt"
)

func main() {
	/*
		//Perulangan For
		for i := 0; i < 5; i++ {
			fmt.Println("Ini perualangan ke ", i)
		}

		//Perulangan While
		l := 1
		for l <= 5 {
			fmt.Println("While ke : ", l)
			l++
		}

		//Perulangan Which biasanya untuk array
		judul := "Semangat Belajar Golang"

		for index, letter := range judul {
			fmt.Println("Index ke ", index, " Huruf : ", string(letter))
		}
	*/

	// Kuis
	judul := "Golang The Best Language"

	for index, huruf := range judul {
		if index%2 == 0 {
			fmt.Println(string(huruf))
		}
	}

	for _, huruf := range judul {
		huruf2 := string(huruf)

		if huruf2 == "a" || huruf2 == "i" || huruf2 == "u" || huruf2 == "e" || huruf2 == "o" {
			fmt.Println(string(huruf))
		}

		switch huruf2 {
		case "a", "i", "u", "e", "o":
			fmt.Println(huruf2)
		}
	}
}
