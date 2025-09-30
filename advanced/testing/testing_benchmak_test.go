package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{10, 20, 30},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	expected := 7
	if result != expected {
		t.Errorf("Add(2, 5) = %d; want %d", result, expected)
	}
}
