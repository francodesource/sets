package sets

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

type Set[T comparable] struct {
	values map[T]struct{}
}

func New[T comparable](elements ...T) Set[T] {
	set := make(map[T]struct{}, len(elements))
	for _, elem := range elements {
		set[elem] = struct{}{}
	}

	return Set[T]{values: set}
}

func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s.values)
}

func (s Set[T]) String() string {
	strs := make([]string, 0, len(s.values))

	for k := range s.values {
		strs = append(strs, fmt.Sprintf("%v", k))
	}

	return "set{" + strings.Join(strs, ", ") + "}"
}

func (s Set[T]) Has(element T) bool {
	_, exists := s.values[element]
	return exists
}

func (s Set[T]) Add(element T) {
	s.values[element] = struct{}{}
}

func (s Set[T]) Remove(element T) {
	delete(s.values, element)
}

func (s Set[T]) Empty() bool {
	return len(s.values) == 0
}

func (s Set[T]) Size() int {
	return len(s.values)
}

func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	for k := range s.values {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

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

func Difference[T comparable](setA, setB Set[T]) Set[T] {
	res := maps.Clone(setA.values)
	for k := range setB.values {
		delete(res, k)
	}
	return Set[T]{values: res}
}
