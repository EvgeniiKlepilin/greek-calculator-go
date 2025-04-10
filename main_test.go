package main

import (
	"reflect"
	"slices"
	"testing"
)

var puzzle [][][]int8 = [][][]int8{
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

func BenchmarkBruteForce(b *testing.B) {
	for b.Loop() {
		bruteForce(puzzle)
	}
}

func TestLayerDials(t *testing.T) {
	layeredDials := layerDials(puzzle)
	expected := [][]int8{
		{12, 10, 6, 16, 10, 7, 10, 8, 1, 4, 9, 2},
		{9, 13, 5, 21, 10, 19, 8, 12, 22, 26, 16, 3},
		{15, 10, 21, 14, 15, 9, 9, 12, 11, 4, 14, 7},
		{15, 8, 8, 11, 3, 11, 6, 6, 10, 7, 7, 14},
	}
	if !reflect.DeepEqual(layeredDials, expected) {
		t.Error("Expected: ", expected, "Actual: ", layeredDials)
	}
}

func BenchmarkLayerDials(b *testing.B) {
	for b.Loop() {
		layerDials(puzzle)
	}
}

func TestCalculateColumn(t *testing.T) {
	result := calculateColumn(puzzle[0], 0)
	expected := 28
	if result != int8(expected) {
		t.Error("Expected: ", expected, "Actual: ", result)
	}
}

func BenchmarkCalculateColumn(b *testing.B) {
	for b.Loop() {
		calculateColumn(puzzle[0], 0)
	}
}

func TestRotateRow(t *testing.T) {
	rotateRow(puzzle, 4, 3)
	expected := []int8{0, 8, 0, 3, 0, 6, 0, 10, 0, 7, 0, 15}
	if !slices.Equal(puzzle[4][3], expected) {
		t.Error("Expected: ", expected, "Actual: ", puzzle[0][0])
	}
}

func BenchmarkRotateRow(b *testing.B) {
	for b.Loop() {
		rotateRow(puzzle, 0, 0)
	}
}

func TestRotateDial(t *testing.T) {
	rotateDial(puzzle, 0)
	expected := [][]int8{
		{10, 7, 16, 8, 7, 8, 8, 3, 4, 12, 2, 5},
		{14, 14, 21, 21, 9, 9, 4, 4, 6, 6, 3, 3},
		{10, 11, 12, 13, 14, 15, 4, 5, 6, 7, 8, 9},
		{14, 14, 11, 14, 11, 14, 11, 11, 14, 11, 14, 11},
	}
	if !reflect.DeepEqual(puzzle[0], expected) {
		t.Error("Expected: ", expected, "Actual: ", puzzle[0])
	}
}

func BenchmarkRotateDial(b *testing.B) {
	for b.Loop() {
		rotateDial(puzzle, 0)
	}
}
