package main

import "testing"

// TestGenerateRandomElements проверяет генерацию слайса случайных чисел.
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{"zero size", 0},
		{"negative size", -1},
		{"one element", 1},
		{"multiple elements", 100},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)

			expectedSize := test.size
			if expectedSize < 0 {
				expectedSize = 0
			}

			if len(data) != expectedSize {
				t.Errorf("expected slice length %d, got %d", expectedSize, len(data))
			}

			for _, value := range data {
				if value <= 0 {
					t.Errorf("expected positive value, got %d", value)
				}
			}
		})
	}
}

// TestMaximum проверяет поиск максимального значения в слайсе.
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{5}, 5},
		{"maximum at beginning", []int{10, 5, 3, 1}, 10},
		{"maximum in middle", []int{1, 5, 20, 7, 3}, 20},
		{"maximum at end", []int{1, 2, 3, 10}, 10},
		{"equal elements", []int{7, 7, 7, 7}, 7},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maximum(test.data)

			if got != test.want {
				t.Errorf("expected %d, got %d", test.want, got)
			}
		})
	}
}
