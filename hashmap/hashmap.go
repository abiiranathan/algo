// implentation of a safe concurrent generic map
package hashmap

import "sync"

// HashMap is a generic implementation of a safe map.
// K is the key and V is the value.
//
// Important for multiple concurrent reads and few writes.
// since it's guarded by sync.RWMutex
type HashMap[K comparable, V any] struct {
	m     map[K]V
	mutex *sync.RWMutex
}

// Instantiates a new AsyncMap.
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		m:     make(map[K]V),
		mutex: &sync.RWMutex{},
	}
}

// Returns V and true if key in map.
func (s *HashMap[K, V]) Get(key K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	v, ok := s.m[key]
	return v, ok
}

// Inserts V in map under the key
func (s *HashMap[K, V]) Set(key K, value V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.m[key] = value
}

// deletes the element with the specified key
func (s *HashMap[K, V]) Delete(key K) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.m, key)
}

// deletes all elements in the map
func (s *HashMap[K, V]) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// delete all elements in the map
	s.m = make(map[K]V)
}

// Returns a slice of the keys in the map
func (s *HashMap[K, V]) Keys() []K {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	keys := make([]K, len(s.m))
	i := 0
	for k := range s.m {
		keys[i] = k
		i++
	}
	return keys
}

// Returns a slice of all the values in the map
// This function copies the underlying data, so it is safe to use
// after the map has been modified.
func (s *HashMap[K, V]) Values() []V {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	values := make([]V, len(s.m))
	i := 0
	for _, v := range s.m {
		values[i] = v
		i++
	}
	return values
}

// deletes all elements in the map
func (s *HashMap[K, V]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.m)
}

// Returns true is map has zero elements
func (s *HashMap[K, V]) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return len(s.m) == 0
}

// Returns true if key in map
func (s *HashMap[K, V]) Contains(key K) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	_, ok := s.m[key]
	return ok
}
