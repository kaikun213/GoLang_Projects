package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for i, _ := range slice {
		slice[i] = make([]uint8, dx)
		for j, _ := range slice[i] {
			slice[i][j] = uint8(i + 1)
		}
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
