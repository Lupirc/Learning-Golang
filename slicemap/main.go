package main

import "fmt"

func main() {
	siswa := []map[string]string{
		{"nama": "Lupi", "nilai": "A"},
		{"nama": "Alpha", "nilai": "B"},
		{"nama": "Beta", "nilai": "C"},
	}

	fmt.Println(siswa)
	fmt.Println("===================")
	for _, students := range siswa {
		fmt.Println("Nama : ", students["nama"], " Nilai : ", students["nilai"])
	}
}
