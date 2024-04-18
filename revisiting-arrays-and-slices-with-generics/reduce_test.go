package main

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	want := 10
	if got := Sum(numbers); got != want {
		t.Errorf("Sum(%v) = %v; want %v", numbers, got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{"Sum all tails", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{2, 4, 6}},
		{"Sum all tails with empty slices", [][]int{{}, {0, 1}, {2, 3, 4}}, []int{0, 1, 7}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumAllTails(tc.input...)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("SumAllTails(%v) = %v; want %v", tc.input, got, tc.expected)
			}
		})
	}
}
