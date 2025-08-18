package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"simple positive", 2, 3, 5},
		{"with zero", 0, 5, 5},
		{"negatives", -2, -3, -5},
		{"mixed", -2, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"multiply by zero", 0, 5, 0},
		{"by negative", -1, 2, -2},
		{"by big numbers", 9000, 3, 27000},
		{"by positive", 10, 18, 180},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ans := multiply(test.a, test.b)
			if ans != test.expected {
				t.Errorf("multiply(%d , %d) = %d; want %d", test.a, test.b, ans, test.expected)
			}
		})
	}
}
