package internal

import (
	"reflect"
	"sort"
	"testing"
)

func TestFactorise(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "Factorise 1",
			input:    1,
			expected: []int{1},
		},
		{
			name:     "Prime number (7)",
			input:    7,
			expected: []int{1, 7},
		},
		{
			name:     "Composite number (12)",
			input:    12,
			expected: []int{1, 2, 3, 4, 6, 12},
		},
		{
			name:     "Perfect square (9)",
			input:    9,
			expected: []int{1, 3, 9},
		},
		{
			name:     "Perfect square (16)",
			input:    16,
			expected: []int{1, 2, 4, 8, 16},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Factorise(tt.input)

			// Sort both slices to ensure order doesn't cause false failures
			sort.Ints(got)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Factorise(%d) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
