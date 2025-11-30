package sets

import "testing"

func TestNewSet(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		expected Set[int]
	}{
		{
			name:     "empty set",
			elements: []int{},
			expected: Set[int]{values: map[int]struct{}{}},
		},
		{
			name:     "set with elements",
			elements: []int{1, 2, 3},
			expected: Set[int]{values: map[int]struct{}{1: {}, 2: {}, 3: {}}},
		},
		{
			name:     "set with duplicate elements",
			elements: []int{1, 2, 2, 3},
			expected: Set[int]{values: map[int]struct{}{1: {}, 2: {}, 3: {}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewSet(tt.elements...)
			if len(result.values) != len(tt.expected.values) {
				t.Errorf("expected set size %d, got %d", len(tt.expected.values), len(result.values))
			}
		})
	}
}

func TestSet_String(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		expected string
	}{
		{
			name:     "empty set",
			elements: []int{},
			expected: "set{}",
		},
		{
			name:     "singleton set",
			elements: []int{1},
			expected: "set{1}",
		},
		{
			name:     "set with elements",
			elements: []int{1, 2, 3},
			expected: "set{1, 2, 3}",
		},
		{
			name:     "set with duplicate elements",
			elements: []int{1, 2, 2, 3},
			expected: "set{1, 2, 3}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.elements...)
			result := set.String()
			if result != tt.expected {
				t.Errorf("expected string %q, got %q", tt.expected, result)
			}
		})
	}
}
