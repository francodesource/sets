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

func TestSet_Iter(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		expected map[int]struct{}
	}{
		{
			name:     "empty set",
			elements: []int{},
			expected: map[int]struct{}{},
		},
		{
			name:     "set with elements",
			elements: []int{1, 2, 3},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.elements...)
			iterSeq := set.Iter()
			result := make(map[int]struct{})
			for val := range iterSeq {
				result[val] = struct{}{}
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected set size %d, got %d", len(tt.expected), len(result))
			}
		})
	}
}

func TestSet_Has(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		check    int
		expected bool
	}{
		{
			name:     "element exists",
			elements: []int{1, 2, 3},
			check:    2,
			expected: true,
		},
		{
			name:     "element does not exist",
			elements: []int{1, 2, 3},
			check:    4,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := NewSet(tt.elements...)
			result := set.Has(tt.check)
			if result != tt.expected {
				t.Errorf("expected Has(%d) to be %v, got %v", tt.check, tt.expected, result)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	if !set.Has(1) {
		t.Errorf("expected set to have element 1 after adding it")
	}
}

func TestSet_Remove(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Remove(2)
	if set.Has(2) {
		t.Errorf("expected set to not have element 2 after removing it")
	}
}

func TestSet_Empty(t *testing.T) {
	set := NewSet[int]()
	if !set.Empty() {
		t.Errorf("expected set to be empty")
	}

	set.Add(1)
	if set.Empty() {
		t.Errorf("expected set to not be empty after adding an element")
	}
}

func TestSet_Size(t *testing.T) {
	set := NewSet(1, 2, 3)
	if set.Size() != 3 {
		t.Errorf("expected set size to be 3, got %d", set.Size())
	}
}
