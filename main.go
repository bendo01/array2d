package main

import "fmt"

func main() {
	var max_baris int
	var max_kolom int
	var arr [3][3]int

	max_baris = 3
	max_kolom = 3

	arr[0][0] = 1
	arr[0][1] = 2
	arr[0][2] = 3
	arr[1][0] = 4
	arr[1][1] = 5
	arr[1][2] = 6
	arr[2][0] = 7
	arr[2][1] = 8
	arr[2][2] = 9


	// ini komentar saya
	for baris := 0; baris < max_baris; baris++ {
		for kolom := 0; kolom < max_kolom; kolom++ {
			fmt.Println(arr[baris][kolom])
		}
	}

}
