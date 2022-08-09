package set

import (
	"github.com/t0mj3dus0r/coffee"
	"github.com/t0mj3dus0r/coffee/iterator"
	"github.com/t0mj3dus0r/coffee/list"
	"github.com/t0mj3dus0r/coffee/streams"
)

func NewHashSet[T comparable]() HashSet[T] {
	return HashSet[T]{
		data: make(map[T]struct{}),
	}
}

type HashSet[T comparable] struct {
	data map[T]struct{}
}

func (h HashSet[T]) Add(t T) bool {
	if h.Contains(t) {
		return false
	}

	h.data[t] = struct{}{}

	return true
}

func (h HashSet[T]) AddAll(t coffee.Collection[T]) bool {

	var result bool

	t.ForEach(func(elem T) {
		setModified := h.Add(elem)
		if setModified {
			result = true
		}
	})

	return result
}

func (h HashSet[T]) Clear() {
	// Optimized by Go compiler. This snippet avoids eventual memory leaks / garbage collecting
	for k := range h.data {
		delete(h.data, k)
	}
}

func (h HashSet[T]) Contains(t T) bool {
	_, exists := h.data[t]

	return exists
}

func (h HashSet[T]) ContainsAll(c coffee.Collection[T]) bool {

	iterator := c.Iterator()

	for iterator.HasNext() {
		if !h.Contains(iterator.Next()) {
			return false
		}
	}

	return true
}

func (h HashSet[T]) IsEmpty() bool {
	return len(h.data) == 0
}

func (h HashSet[T]) Size() int {
	return len(h.data)
}

func (h HashSet[T]) ToArray() []T {

	size := len(h.data)

	result := make([]T, size)

	var index int

	for k := range h.data {
		result[index] = k
		index++
	}

	return result
}

func (h HashSet[T]) Remove(t T) bool {
	if !h.Contains(t) {
		return false
	}

	delete(h.data, t)

	return true
}

func (h HashSet[T]) RemoveIf(filter coffee.Predicate[T]) bool {

	var result bool

	h.ForEach(func(elem T) {
		if filter.Test(elem) {
			h.Remove(elem)
			result = true
		}
	})

	return result
}

func (h HashSet[T]) RemoveAll(c coffee.Collection[T]) bool {

	var result bool

	c.ForEach(func(elem T) {
		if c.Contains(elem) {
			h.Remove(elem)
			result = true
		}
	})

	return result
}

func (h HashSet[T]) RetainAll(c coffee.Collection[T]) bool {

	existingElements := list.NewArrayList[T]()

	c.ForEach(func(elem T) {
		if h.Contains(elem) {
			existingElements.Add(elem)
		}
	})

	result := existingElements.Size() != h.Size()

	h.Clear()

	h.AddAll(existingElements)

	return result
}

func (h HashSet[T]) ForEach(action coffee.Consumer[T]) {
	for k := range h.data {
		action(k)
	}
}

func (h HashSet[T]) Iterator() coffee.Iterator[T] {
	return iterator.List[T](
		list.FromArray(
			h.ToArray(),
		),
	)
}

func (h HashSet[T]) Stream() coffee.Stream[T] {
	return streams.FromArray(
		h.ToArray(),
	)
}
