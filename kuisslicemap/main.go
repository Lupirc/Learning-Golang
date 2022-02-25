package main

import "fmt"

func main() {
	//Menghitung rata-rata
	score := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var jml int

	for _, nilai := range score {
		//jml += nilai
		jml = jml + nilai
	}

	rata2 := float64(jml) / float64(len(score))
	fmt.Println(rata2)
	fmt.Println("===================")

	//Slice goodScore
	var goodScore []int

	for _, nilai := range score {
		if nilai >= 90 {
			goodScore = append(goodScore, nilai)
		}
	}

	fmt.Println(goodScore)
}
