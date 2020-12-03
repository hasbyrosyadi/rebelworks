package main

import (
	"fmt"
	"sort"
)

func SearchParent(input int) int {

	result := 0
	if input%3 == 0 {
		result = input / 3
	} else if input%3 == 1 {
		result = (input - 1) / 3
	} else if input%3 == 2 {
		result = (input + 1) / 3
	}

	return result
}

func GetPersentage(jumlahAnak []int) []float64 {

	var persentage []float64
	for _, v := range jumlahAnak {
		if v >= 11 && v <= 15 {
			persentage = append(persentage, 0.1)
		} else if v >= 6 && v <= 10 {
			persentage = append(persentage, 0.07)
		} else if v >= 1 && v <= 5 {
			persentage = append(persentage, 0.05)
		}
	}
	return persentage
}

func Tunjangan(gajiPokok float64, jumlahAnak []int) float64 {
	var tunjangan float64 = 0
	if len(jumlahAnak) == 1 {
		tunjangan = 0.1 * gajiPokok
		return tunjangan
	}

	// mengurutkan dari umur dari yang terbesar
	sort.Sort(sort.Reverse(sort.IntSlice(jumlahAnak)))

	// mengambil 2 terbesar
	persentage := GetPersentage(jumlahAnak[:2])
	for _, v := range persentage {
		tunjangan += v * gajiPokok
	}

	return tunjangan
}

func main() {
	// Node
	fmt.Println(SearchParent(2100))

	// Coding Test
	gajiPokok := 3500000
	jumlahAnak := []int{14, 3, 5, 15}
	fmt.Println(Tunjangan(float64(gajiPokok), jumlahAnak))
}
