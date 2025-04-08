package main

import "fmt"

func main() {
	puzzle := [][][]int8{
		{
			{5, 10, 7, 16, 8, 7, 8, 8, 3, 4, 12, 2},
			{3, 14, 14, 21, 21, 9, 9, 4, 4, 6, 6, 3},
			{9, 10, 11, 12, 13, 14, 15, 4, 5, 6, 7, 8},
			{11, 14, 14, 11, 14, 11, 14, 11, 11, 14, 11, 14},
		},
		{
			{12, 0, 6, 0, 10, 0, 10, 0, 1, 0, 9, 0},
			{2, 13, 9, 0, 17, 19, 3, 12, 3, 26, 6, 0},
			{6, 0, 14, 12, 3, 8, 9, 0, 9, 20, 12, 3},
			{7, 14, 11, 0, 8, 0, 16, 2, 7, 0, 9, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{9, 0, 5, 0, 10, 0, 8, 0, 22, 0, 16, 0},
			{12, 0, 21, 6, 15, 4, 9, 18, 11, 26, 14, 1},
			{7, 8, 9, 13, 9, 7, 13, 21, 17, 4, 5, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{15, 0, 0, 14, 0, 9, 0, 12, 0, 4, 0, 7},
			{6, 0, 11, 11, 6, 11, 0, 6, 17, 7, 3, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{15, 0, 8, 0, 3, 0, 6, 0, 10, 0, 7, 0},
		},
	}

	bruteForce(puzzle)
}

func bruteForce(s [][][]int8) {
	for range 12 {
		for range 12 {
			for range 12 {
				for range 12 {
					layerCalculateAndCheck(s)
					rotateDial(s, 4)
				}
				rotateDial(s, 3)
			}
			rotateDial(s, 2)
		}
		rotateDial(s, 1)
	}
	rotateDial(s, 0)
}

func layerCalculateAndCheck(s [][][]int8) {
	layeredDials := layerDials(s)
	_, isSolved := calculateColumns(layeredDials, false)
	if isSolved {
		fmt.Println("Solution: ", layeredDials)
	}
}

func rotateDial(s [][][]int8, dial int8) {
	for rowIndex := range s[dial] {
		rotateRow(s, dial, int8(rowIndex))
	}
}

func rotateRow(s [][][]int8, dial int8, row int8) {
	s[dial][row] = append(s[dial][row][1:], s[dial][row][0])
}

func calculateColumn(s [][]int8, column int8) (sum int8) {
	for rowIndex := range s {
		sum += s[rowIndex][column]
	}
	return
}

func calculateColumns(s [][]int8, calculateAll bool) (sums []int8, ok bool) {
	for columnIndex := range s[0] {
		sum := calculateColumn(s, int8(columnIndex))
		if sum != 42 {
			ok = false
			if !calculateAll {
				return
			}
		}
		sums = append(sums, sum)
	}
	ok = true
	return
}

func layerDials(s [][][]int8) (layeredDials [][]int8) {
	layeredDials = make([][]int8, 4)
	for i := range layeredDials {
		layeredDials[i] = make([]int8, 12)
	}
	for dial := len(s) - 1; dial >= 0; dial-- {
		for rowIndex := range s[dial] {
			for columnIndex, value := range s[dial][rowIndex] {
				if layeredDials[rowIndex][columnIndex] == 0 {
					layeredDials[rowIndex][columnIndex] = value
				}
			}
		}
	}
	return
}
