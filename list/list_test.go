package list

import (
	"fmt"
	"testing"
)

func TestNewList(t *testing.T) {
	// Initialize with only length
	l := New[int](10)
	if l.Len() != 10 {
		t.Fatalf("expected list.Len() to be 10, got %d", l.Len())
	}

	// Initialize with length and capacity
	l2 := New[int](0, 10)
	if l2.Len() != 0 {
		t.Fatalf("invalid initial length")
	}

	// Initialize with no arguments
	// Underlying slice is not pre-allocated
	l3 := New[int]()
	if l3.Len() != 0 {
		t.Fatalf("invalid initial length")
	}
}

func TestListAppend(t *testing.T) {
	l := New[string](0, 5)
	l.Append("name", "age")

	if l.Len() != 2 {
		t.Errorf("expected l.Append to increase length by 2")
	}

	index := l.Index(func(val string) bool {
		return val == "name"
	})

	if index != 0 {
		t.Errorf("expected index of name to be 0, got %d", index)
	}

	index = l.Index(func(val string) bool {
		return val == "age"
	})

	if index != 1 {
		t.Errorf("expected index of age to be 1, got %d", index)
	}

	// test that out-of-range search returns index of -1
	index = l.Index(func(val string) bool {
		return val == "out of range"
	})

	if index != -1 {
		t.Errorf("out of bounds search should return -1, got :%d", index)
	}
}

func TestInsert(t *testing.T) {
	l := New[int](4)

	tt := []struct {
		index int
		value int
	}{
		{2, 30},
		{3, 40},
		{0, 10},
		{1, 0},
	}

	for _, test := range tt {
		l.Insert(test.index, test.value)

		if index := l.Index(func(val int) bool {
			return test.value == val
		}); index != test.index {
			t.Errorf("l.Insert %d at wrong index. Expected %d, got %d", test.value, test.index, index)
		}
	}
}

func TestSlice(t *testing.T) {
	l := New[int](0, 5)
	l.Append(1, 2, 3, 4, 5)

	s := l.Slice(2, l.Len())
	if s.Len() != 3 {
		t.Errorf("slicing returned a new slice of incorrect len, got %d, expected %d", s.Len(), 3)
	}
}

func TestRemove(t *testing.T) {
	l := New[int](0, 5)
	l.Append(1, 2, 3, 4, 5)

	l.Remove(2)
	l.Remove(3)

	if l.Len() != 3 {
		t.Errorf("l.Len() after remove is incorrect. Expected 3, got %d", l.Len())
	}

	Map(l, func(v int) int {
		return v * 10
	}).ForEach(func(index, val int) {
	})
}

func TestFilter(t *testing.T) {
	nums := New[int](0, 10)
	nums.Append(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	evens := nums.Filter(func(val int) bool {
		return val%2 == 0
	})

	evens.ForEach(func(index, val int) {
		if val%2 != 0 {
			t.Errorf("not an even number, %d", val)
		}
	})

	if evens.Len() != 5 {
		t.Errorf("expected length of 5, got %d", evens.Len())
	}

	evens.ForEachPtr(func(index int, val *int) {
		if *val%2 != 0 {
			t.Errorf("not an even number, %d", val)
		}
	})
}

func TestCopy(t *testing.T) {
	s := New[int](0, 1000)

	for i := 0; i < 1000; i++ {
		s.Append(i)
	}

	newList := s.Clone()

	if s.Len() != newList.Len() {
		t.Errorf("mismatch in len")
	}

	if s.Get(0) != newList.Get(0) {
		t.Error("first element mismatch")
	}

	if s.Get(s.Len()-1) != newList.Get(newList.Len()-1) {
		t.Error("last element mismatch")
	}
}

func square(_, val int) int {
	return val * val
}

func BenchmarkNormalSlice(b *testing.B) {
	s := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		s[i] = i
	}

	for i, val := range s {
		sq := square(i, val)
		fmt.Println(sq)
	}

	// 1000000000	         0.002287 ns/op	       0 B/op	       0 allocs/op

}

func BenchmarkForEachLoop(b *testing.B) {
	s := New[int](0, 1000)

	for i := 0; i < 1000; i++ {
		s.Append(i)
	}

	s.ForEach(func(i, val int) {
		sq := square(i, val)
		fmt.Println(sq)
	})

	// 1000000000	         0.002149 ns/op	       0 B/op	       0 allocs/op
	// It's clear that there is no significant slow down with generic code.
}
