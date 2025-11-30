package sets

import (
	"regexp"
	"testing"
)

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
			result := New(tt.elements...)
			if len(result.values) != len(tt.expected.values) {
				t.Errorf("expected set size %d, got %d", len(tt.expected.values), len(result.values))
			}
		})
	}
}

func TestWithCapacity(t *testing.T) {
	set := WithCapacity[int](3)
	if len(set.values) != 0 {
		t.Errorf("expected empty set, got size %d", len(set.values))
	}
}

func TestSet_String(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		expected *regexp.Regexp
	}{
		{
			name:     "empty set",
			elements: []int{},
			expected: regexp.MustCompile(`^set\{}$`),
		},
		{
			name:     "singleton set",
			elements: []int{1},
			expected: regexp.MustCompile(`^set\{1}$`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New(tt.elements...)
			result := set.String()
			if !tt.expected.MatchString(result) {
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
			set := New(tt.elements...)
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
			set := New(tt.elements...)
			result := set.Has(tt.check)
			if result != tt.expected {
				t.Errorf("expected Has(%d) to be %v, got %v", tt.check, tt.expected, result)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	set := New[int]()
	set.Add(1)
	if !set.Has(1) {
		t.Errorf("expected set to have element 1 after adding it")
	}
}

func TestSet_Remove(t *testing.T) {
	set := New(1, 2, 3)
	set.Remove(2)
	if set.Has(2) {
		t.Errorf("expected set to not have element 2 after removing it")
	}
}

func TestSet_Empty(t *testing.T) {
	set := New[int]()
	if !set.Empty() {
		t.Errorf("expected set to be empty")
	}

	set.Add(1)
	if set.Empty() {
		t.Errorf("expected set to not be empty after adding an element")
	}
}

func TestSet_Size(t *testing.T) {
	set := New(1, 2, 3)
	if set.Size() != 3 {
		t.Errorf("expected set size to be 3, got %d", set.Size())
	}
}

func TestSet_Equals(t *testing.T) {
	tests := []struct {
		name     string
		set1     Set[int]
		set2     Set[int]
		expected bool
	}{
		{
			name:     "equal sets",
			set1:     New(1, 2, 3),
			set2:     New(1, 2, 3),
			expected: true,
		},
		{
			name:     "unequal sets",
			set1:     New(1, 2, 3),
			set2:     New(1, 2, 4),
			expected: false,
		},
		{
			name:     "different sizes",
			set1:     New(1, 2),
			set2:     New(1, 2, 3),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.set1.Equals(tt.set2)
			if result != tt.expected {
				t.Errorf("expected Equals to be %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSet_IsSubsetOf(t *testing.T) {
	tests := []struct {
		name     string
		subset   Set[int]
		superset Set[int]
		expected bool
	}{
		{
			name:     "is subset",
			subset:   New(1, 2),
			superset: New(1, 2, 3),
			expected: true,
		},
		{
			name:     "is not subset",
			subset:   New(1, 4),
			superset: New(1, 2, 3),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.subset.IsSubsetOf(tt.superset)
			if result != tt.expected {
				t.Errorf("expected IsSubsetOf to be %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name     string
		sets     []Set[int]
		expected []int
	}{
		{
			name:     "union of two sets",
			sets:     []Set[int]{New(1, 2), New(2, 3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "union with empty set",
			sets:     []Set[int]{New(1, 2), New[int]()},
			expected: []int{1, 2},
		},
		{
			name:     "union of multiple sets",
			sets:     []Set[int]{New(1), New(2), New(3)},
			expected: []int{1, 2, 3},
		},
		{
			name:     "union of no sets",
			sets:     []Set[int]{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Union(tt.sets...)
			expectedSet := New(tt.expected...)
			if !result.Equals(expectedSet) {
				t.Errorf("expected union to be %v, got %v", expectedSet, result)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		sets     []Set[int]
		expected []int
	}{
		{
			name:     "intersection of two sets",
			sets:     []Set[int]{New(1, 2), New(2, 3)},
			expected: []int{2},
		},
		{
			name:     "intersection with empty set",
			sets:     []Set[int]{New(1, 2), New[int]()},
			expected: []int{},
		},
		{
			name:     "intersection of multiple sets",
			sets:     []Set[int]{New(1, 2, 3), New(2, 3, 4), New(3, 4, 5)},
			expected: []int{3},
		},
		{
			name:     "intersection of no sets",
			sets:     []Set[int]{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Intersection(tt.sets...)
			expectedSet := New(tt.expected...)
			if !result.Equals(expectedSet) {
				t.Errorf("expected intersection to be %v, got %v", expectedSet, result)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		setA     Set[int]
		setB     Set[int]
		expected []int
	}{
		{
			name:     "difference of two sets",
			setA:     New(1, 2, 3),
			setB:     New(2, 3),
			expected: []int{1},
		},
		{
			name:     "difference with empty set",
			setA:     New(1, 2),
			setB:     New[int](),
			expected: []int{1, 2},
		},
		{
			name:     "difference resulting in empty set",
			setA:     New(1, 2),
			setB:     New(1, 2),
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.setA, tt.setB)
			expectedSet := New(tt.expected...)
			if !result.Equals(expectedSet) {
				t.Errorf("expected difference to be %v, got %v", expectedSet, result)
			}
		})
	}
}
