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

func NewSet[T comparable](elements ...T) Set[T] {
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
