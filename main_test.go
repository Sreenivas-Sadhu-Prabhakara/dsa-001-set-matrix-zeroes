package main

import (
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	cases := []struct {
		name string
		in   [][]int
		want [][]int
	}{
		{"center zero", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"first row and col", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
		{"no zeroes", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"all zero", [][]int{{0}}, [][]int{{0}}},
		{"empty", [][]int{}, [][]int{}},
	}
	for _, c := range cases {
		SetZeroes(c.in)
		if !reflect.DeepEqual(c.in, c.want) {
			t.Errorf("%s: got %v want %v", c.name, c.in, c.want)
		}
	}
}
