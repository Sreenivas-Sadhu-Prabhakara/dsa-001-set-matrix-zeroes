package main

import "fmt"

// SetZeroes zeroes every row and column that contains a zero, in place.
// The first row and first column double as the marker store, so no extra
// row/column arrays are allocated.
func SetZeroes(m [][]int) {
	if len(m) == 0 || len(m[0]) == 0 {
		return
	}
	rows, cols := len(m), len(m[0])
	firstRowZero, firstColZero := false, false
	for c := 0; c < cols; c++ {
		if m[0][c] == 0 {
			firstRowZero = true
		}
	}
	for r := 0; r < rows; r++ {
		if m[r][0] == 0 {
			firstColZero = true
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if m[r][c] == 0 {
				m[r][0] = 0
				m[0][c] = 0
			}
		}
	}
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if m[r][0] == 0 || m[0][c] == 0 {
				m[r][c] = 0
			}
		}
	}
	if firstRowZero {
		for c := 0; c < cols; c++ {
			m[0][c] = 0
		}
	}
	if firstColZero {
		for r := 0; r < rows; r++ {
			m[r][0] = 0
		}
	}
}

func main() {
	m := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	SetZeroes(m)
	fmt.Println(m)
}
