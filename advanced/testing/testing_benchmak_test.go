package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100) // Random integers between 0 and 99
	}

	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
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

func TestAddSubTests(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{10, 20, 30},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	expected := 7
	if result != expected {
		t.Errorf("Add(2, 5) = %d; want %d", result, expected)
	}
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 10
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("GenerateRandomSlice(%d) length = %d; want %d", size, len(slice), size)
	}

	for _, v := range slice {
		if v < 0 || v >= 100 {
			t.Errorf("GenerateRandomSlice(%d) contains out-of-range value: %d", size, v)
		}
	}
}

func TestSumSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5} // Fixed slice for consistent testing
	result := SumSlice(slice)
	expected := len(slice) * (slice[0] + slice[len(slice)-1]) / 2 // Sum of first n natural numbers formula

	if result != expected {
		t.Errorf("SumSlice(%v) = %d; want %d", slice, result, expected)
	}
}

// Benchmarking - measures the performance of a code
// - How long it takes to execute a piece of code
func BenchmarkAddSmallInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1000, 2000)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1782635343, 28273534250)
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomSlice(1000) // Benchmarking with a slice of size 1000
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		SumSlice(slice)
	}
}
