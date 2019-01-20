package utils_test

import (
	"noxes/backend/utils"
	"testing"
)

func TestNilSlicesAreEqual(t *testing.T) {
	var slice1 []int = nil
	var slice2 []int = nil

	if !utils.SlicesAreEqual(slice1, slice2) {
		t.Fatalf("Nil slices are not equal")
	}
}

func TestNilSliceDoesNotEqualNonNilSlice(t *testing.T) {
	var slice1 []int = nil
	var slice2 = []int{1, 2, 3}

	if utils.SlicesAreEqual(slice1, slice2) {
		t.Fatalf("Nil slice should not equal non-nil slice")
	}
}

func TestDifferentLengthSlicesAreNotEqual(t *testing.T) {
	var slice1 = make([]int, 0)
	var slice2 = []int{1, 2, 3}

	if utils.SlicesAreEqual(slice1, slice2) {
		t.Fatalf("Different length slices should not be equal")
	}
}

func TestSlicesWithSameLengthDataAreEqual(t *testing.T) {
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 3}

	if !utils.SlicesAreEqual(slice1, slice2) {
		t.Fatalf("Slices with same length and data should be equal")
	}
}

func TestSlicesWithSameLengthButDifferentDataAreNotEqual(t *testing.T) {
	var slice1 = []int{1, 2, 6}
	var slice2 = []int{1, 2, 3}

	if utils.SlicesAreEqual(slice1, slice2) {
		t.Fatalf("Slices with same length but different data should not equal")
	}
}
