package sets

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

// Set is a generic set data structure for comparable types, without duplicates.
// Order of elements is not guaranteed.
type Set[T comparable] struct {
	values map[T]struct{}
}

// New creates a new Set with the given elements.
func New[T comparable](elements ...T) Set[T] {
	set := make(map[T]struct{}, len(elements))
	for _, elem := range elements {
		set[elem] = struct{}{}
	}

	return Set[T]{values: set}
}

// WithCapacity creates a new Set with the specified initial capacity.
func WithCapacity[T comparable](capacity int) Set[T] {
	return Set[T]{values: make(map[T]struct{}, capacity)}
}

// Iter returns an iterator over the elements of the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s.values)
}

// String returns a string representation of the set.
func (s Set[T]) String() string {
	strs := make([]string, 0, len(s.values))

	for k := range s.values {
		strs = append(strs, fmt.Sprintf("%v", k))
	}

	return "set{" + strings.Join(strs, ", ") + "}"
}

// Has checks if the element is in the set.
func (s Set[T]) Has(element T) bool {
	_, exists := s.values[element]
	return exists
}

// Add adds an element to the set.
func (s Set[T]) Add(element T) {
	s.values[element] = struct{}{}
}

// Remove removes an element from the set.
func (s Set[T]) Remove(element T) {
	delete(s.values, element)
}

// Empty checks if the set is empty.
func (s Set[T]) Empty() bool {
	return len(s.values) == 0
}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s.values)
}

// IsSubsetOf checks if the set is a subset of another set.
func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	for k := range s.values {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

// Equals checks if the set is equal to another set.
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for k := range s.values {
		if !other.Has(k) {
			return false
		}
	}

	return true
}

// Copy creates a shallow copy of the set.
func Copy[T comparable](s Set[T]) Set[T] {
	return Set[T]{values: maps.Clone(s.values)}
}

// Union returns a new set that is the union of the provided sets.
func Union[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return Set[T]{values: make(map[T]struct{})}
	}

	res := maps.Clone(sets[0].values)
	for _, set := range sets[1:] {
		maps.Copy(res, set.values)
	}
	return Set[T]{values: res}
}

// Intersection returns a new set that is the intersection of the provided sets.
func Intersection[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return Set[T]{values: make(map[T]struct{})}
	}

	res := maps.Clone(sets[0].values)
	for _, set := range sets[1:] {
		maps.DeleteFunc(res, func(key T, _ struct{}) bool {
			_, exists := set.values[key]
			return !exists
		})
	}
	return Set[T]{values: res}
}

// Difference returns a new set that is the difference of setA and setB (elements in setA not in setB).
func Difference[T comparable](setA, setB Set[T]) Set[T] {
	res := maps.Clone(setA.values)
	for k := range setB.values {
		delete(res, k)
	}
	return Set[T]{values: res}
}
