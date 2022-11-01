package list

// golang generic slice data structure.
type List[T any] interface {
	// Returns the length of the underlying slice
	Len() int

	/*
		Returns the index given a predicate that is called for each value
		in the list. The predicate function should return true for the value
		you want. This avoid strict == checks that would only work for native
		data types.
			index := list.Index(func(val int) bool {
				return val == 20
			})

		Big O(n) in worst-case scenario.
	*/
	Index(cmp func(val T) bool) int

	/*
		Append elements into the list.
		Append will resize the underlying slice if necessary.
	*/
	Append(vals ...T)

	/*
		Insert inserts a value at index in the under lying slice.
		index should not be out-of-bounds, otherwise this method will panic.
	*/
	Insert(index int, val T)

	// At returns element at the given index
	// Does not do out-of-bounds check on the index and will panic if index
	// is not valid.
	Get(index int) T

	/* Creates a slice expression from start to end and wraps it in a new List.
	Constraints:
	- start and end must be in the range capacity of list(inclusive)
	- start <= end
	*/
	Slice(start int, end int) List[T]

	/*copies this list into a new list of the same length and capacity*/
	Clone() List[T]

	/*Removes the item at index. index should be within bounds otherwise this method will panic
	out-of-bounds. Remove if successful will change the length of the underling slice.
	*/
	Remove(index int)

	// ForEach iterates and passes the index and value to the callback.
	ForEach(callback func(index int, val T))

	// given a callback, ForEach iterates and passes the index and pointer to value to the callback.
	ForEachPtr(callback func(index int, val *T))

	// Filter items in list given a predicate.
	// Filter does not mutate the underlying array.
	Filter(predicate func(val T) bool) List[T]
}

// implements List interface
type list[T any] struct {
	s []T // the underlying slice
}

// Map transforms list l of type T into a list of V
func Map[T any, V any](l List[T], callback func(v T) V) List[V] {
	new_slice := make([]V, l.Len())
	l.ForEach(func(index int, val T) {
		new_value := callback(val)
		new_slice[index] = new_value
	})
	return NewListFromSlice(&new_slice)
}

/*
Takes the same arguments as make, except that if no size is provided,
the undelying slice is not pre-allocated.
*/
func New[T any](size ...int) List[T] {
	if len(size) == 1 {
		return &list[T]{
			s: make([]T, size[0]),
		}
	} else if len(size) == 2 {
		return &list[T]{
			s: make([]T, size[0], size[1]),
		}
	} else {
		return &list[T]{
			s: []T{},
		}
	}

}

func NewListFromSlice[T any](l *[]T) List[T] {
	ret := New[T](len(*l), cap(*l))
	for i, v := range *l {
		ret.Insert(i, v)
	}
	return ret
}

func (l *list[T]) Len() (length int) {
	return len(l.s)
}

func (l *list[T]) Index(cmp func(val T) bool) int {
	for index, value := range l.s {
		if cmp(value) {
			return index
		}
	}
	return -1
}

func (l *list[T]) Append(vals ...T) {
	l.s = append(l.s, vals...)
}

func (l *list[T]) Insert(index int, val T) {
	l.s[index] = val
}

func (l *list[T]) Slice(start int, end int) List[T] {
	expr := l.s[start:end]
	return NewListFromSlice(&expr)
}

func (l *list[T]) Clone() List[T] {
	return NewListFromSlice(&l.s)
}

func (l *list[T]) Remove(index int) {
	l.s = append(l.s[:index], l.s[index+1:]...)
}

func (l *list[T]) ForEach(callback func(index int, val T)) {
	for i, v := range l.s {
		callback(i, v)
	}
}

func (l *list[T]) ForEachPtr(callback func(index int, val *T)) {
	for i := range l.s {
		callback(i, &l.s[i])
	}
}

func (l *list[T]) Filter(predicate func(val T) bool) List[T] {
	ret := New[T](0, l.Len())
	for _, v := range l.s {
		if predicate(v) {
			ret.Append(v)
		}
	}
	return ret
}

func (l *list[T]) Get(index int) T {
	return l.s[index]
}
