package set

// Generic Set data structure implemention.
type Set[T comparable] struct {
	hash map[T]struct{}
}

// Create a new set
func New[T comparable](initial ...T) *Set[T] {
	s := &Set[T]{hash: make(map[T]struct{})}

	for _, v := range initial {
		s.Insert(v)
	}
	return s
}

// Find the difference between two sets
func (s *Set[T]) Difference(set *Set[T]) *Set[T] {
	n := make(map[T]struct{})

	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			n[k] = struct{}{}
		}
	}
	return &Set[T]{n}
}

// Call f for each item in the set
func (s *Set[T]) ForEach(f func(elem T)) {
	for k := range s.hash {
		f(k)
	}
}

// Test to see whether or not the element is in the set
func (s *Set[T]) Has(element T) bool {
	_, exists := s.hash[element]
	return exists
}

// Add an element to the set
func (s *Set[T]) Insert(element T) {
	s.hash[element] = struct{}{}
}

// Find the intersection of two sets
func (s *Set[T]) Intersection(set *Set[T]) *Set[T] {
	n := make(map[T]struct{})

	for k := range s.hash {
		if _, exists := set.hash[k]; exists {
			n[k] = struct{}{}
		}
	}
	return &Set[T]{n}
}

// Return the number of items in the set
func (s *Set[T]) Len() int {
	return len(s.hash)
}

// Test whether or not this set is a proper subset of "set"
func (s *Set[T]) ProperSubsetOf(set *Set[T]) bool {
	return s.SubsetOf(set) && s.Len() < set.Len()
}

// Remove an element from the set
func (s *Set[T]) Remove(element T) {
	delete(s.hash, element)
}

// Test whether or not this set is a subset of "set"
func (s *Set[T]) SubsetOf(set *Set[T]) bool {
	if s.Len() > set.Len() {
		return false
	}

	for k := range s.hash {
		if _, exists := set.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Find the union of two sets
func (s *Set[T]) Union(set *Set[T]) *Set[T] {
	u := make(map[T]struct{})

	for k := range s.hash {
		u[k] = struct{}{}
	}

	for k := range set.hash {
		u[k] = struct{}{}
	}

	return &Set[T]{u}
}
