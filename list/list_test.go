package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	l := New()

	require.Nil(t, l.first)
	require.Nil(t, l.last)
	require.Equal(t, 0, l.size)
}

func TestSize(t *testing.T) {
	l := New()

	require.Equal(t, 0, Size(l))
	require.Equal(t, true, Empty(l))

	// HACK
	l.size = 1

	require.Equal(t, 1, Size(l))
	require.Equal(t, false, Empty(l))
}

func TestPushFront(t *testing.T) {
	l := New()

	require.Nil(t, l.first)
	require.Nil(t, l.last)
	require.Equal(t, 0, Size(l))

	PushFront(l, 42)

	require.Equal(t, 42, l.first.value)
	require.Equal(t, 42, l.last.value)
	require.Equal(t, 1, Size(l))

	PushFront(l, 45)

	require.Equal(t, 45, l.first.value)
	require.Equal(t, 42, l.last.value)
	require.Equal(t, 2, Size(l))

	require.Equal(t, l.last, l.first.next)
}

func TestPushBack(t *testing.T) {
	l := New()

	require.Nil(t, l.first)
	require.Nil(t, l.last)
	require.Equal(t, 0, Size(l))

	PushBack(l, 42)

	require.Equal(t, 42, l.first.value)
	require.Equal(t, 42, l.last.value)
	require.Equal(t, 1, Size(l))

	PushBack(l, 45)

	require.Equal(t, 42, l.first.value)
	require.Equal(t, 45, l.last.value)
	require.Equal(t, 2, Size(l))

	require.Equal(t, l.last, l.first.next)
}

func TestAt(t *testing.T) {
	l := New()
	var (
		res int
		ok  bool
	)

	res, ok = At(l, 0)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)

	res, ok = At(l, 0)
	require.Equal(t, true, ok)
	require.Equal(t, 1, res)

	res, ok = At(l, 1)
	require.Equal(t, true, ok)
	require.Equal(t, 2, res)

	res, ok = At(l, 2)
	require.Equal(t, true, ok)
	require.Equal(t, 3, res)

	require.Equal(t, 1, l.first.value)
	require.Equal(t, 3, l.last.value)

	res, ok = At(l, -1)
	require.Equal(t, true, ok)
	require.Equal(t, 2, res)

	res, ok = At(l, -2)
	require.Equal(t, true, ok)
	require.Equal(t, 1, res)

	res, ok = At(l, -3)
	require.Equal(t, false, ok)
}

func TestRemove(t *testing.T) {
	l := New()
	var ok bool

	ok = Remove(l, 0)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)
	PushBack(l, 4)

	require.Equal(t, 1, l.first.value)
	require.Equal(t, 4, l.last.value)
	require.Equal(t, 4, Size(l))

	// Remove the second element, the list should be [1, 3, 4]
	ok = Remove(l, 1)
	require.Equal(t, true, ok)
	require.Equal(t, 1, l.first.value)
	require.Equal(t, 4, l.last.value)
	require.Equal(t, 3, Size(l))

	// Remove the first element, the list should be [3, 4]
	ok = Remove(l, 0)
	require.Equal(t, true, ok)
	require.Equal(t, 3, l.first.value)
	require.Equal(t, 4, l.last.value)
	require.Equal(t, 2, Size(l))

	// Remove the last element, the list should be [3]
	ok = Remove(l, 1)
	require.Equal(t, true, ok)
	require.Equal(t, 3, l.first.value)
	require.Equal(t, 3, l.last.value)
	require.Equal(t, l.first, l.last)
	require.Equal(t, 1, Size(l))

	// Remove the only element, the list should be empty
	ok = Remove(l, 0)
	require.Equal(t, true, ok)
	require.Nil(t, l.first)
	require.Nil(t, l.last)
	require.Equal(t, 0, Size(l))
}

func TestPopBack(t *testing.T) {
	var (
		res int
		ok  bool
	)

	l := New()

	_, ok = PopBack(l)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)
	PushBack(l, 4)
	require.Equal(t, 4, Size(l))

	res, ok = PopBack(l)
	require.Equal(t, true, ok)
	require.Equal(t, 4, res)
	require.Equal(t, 3, Size(l))

	res, ok = PopBack(l)
	require.Equal(t, true, ok)
	require.Equal(t, 3, res)
	require.Equal(t, 2, Size(l))

	res, ok = PopBack(l)
	require.Equal(t, true, ok)
	require.Equal(t, 2, res)
	require.Equal(t, 1, Size(l))

	res, ok = PopBack(l)
	require.Equal(t, true, ok)
	require.Equal(t, 1, res)
	require.Equal(t, 0, Size(l))
}

func TestPopFront(t *testing.T) {
	var (
		res int
		ok  bool
	)

	l := New()

	_, ok = PopFront(l)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)
	PushBack(l, 4)
	require.Equal(t, 4, Size(l))

	res, ok = PopFront(l)
	require.Equal(t, true, ok)
	require.Equal(t, 1, res)
	require.Equal(t, 3, Size(l))

	res, ok = PopFront(l)
	require.Equal(t, true, ok)
	require.Equal(t, 2, res)
	require.Equal(t, 2, Size(l))

	res, ok = PopFront(l)
	require.Equal(t, true, ok)
	require.Equal(t, 3, res)
	require.Equal(t, 1, Size(l))

	res, ok = PopFront(l)
	require.Equal(t, true, ok)
	require.Equal(t, 4, res)
	require.Equal(t, 0, Size(l))
}

func TestFront(t *testing.T) {
	l := New()

	var (
		res int
		ok  bool
	)

	_, ok = Front(l)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)
	PushBack(l, 4)

	res, ok = Front(l)
	require.Equal(t, true, ok)
	require.Equal(t, 1, res)
}

func TestBack(t *testing.T) {
	l := New()

	var (
		res int
		ok  bool
	)

	_, ok = Back(l)
	require.Equal(t, false, ok)

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)
	PushBack(l, 4)

	res, ok = Back(l)
	require.Equal(t, true, ok)
	require.Equal(t, 4, res)
}

func TestInsert(t *testing.T) {
	l := New()
	require.Nil(t, l.first)
	require.Nil(t, l.last)
	require.Equal(t, 0, Size(l))

	var ok bool

	ok = Insert(l, 0, 4)
	require.Equal(t, true, ok)
	require.Equal(t, 4, l.first.value)
	require.Equal(t, l.first, l.last)
	require.Equal(t, 1, Size(l))

	ok = Insert(l, 0, 5)
	require.Equal(t, true, ok)
	require.Equal(t, 5, l.first.value)
	require.Equal(t, 4, l.last.value)
	require.Equal(t, l.last, l.first.next)
	require.Equal(t, 2, Size(l))

	ok = Insert(l, 1, 6)
	require.Equal(t, true, ok)
	require.Equal(t, 5, l.first.value)
	require.Equal(t, 6, l.first.next.value)
	require.Equal(t, 4, l.last.value)
	require.Equal(t, 3, Size(l))

	ok = Insert(l, 3, 7)
	require.Equal(t, true, ok)
	require.Equal(t, 5, l.first.value)
	require.Equal(t, 6, l.first.next.value)
	require.Equal(t, 4, l.first.next.next.value)
	require.Equal(t, 7, l.last.value)
	require.Equal(t, 4, Size(l))

	ok = Insert(l, 5, 8)
	require.Equal(t, false, ok)

	ok = Insert(l, -1, 8)
	require.Equal(t, true, ok)
	require.Equal(t, 5, l.first.value)
	require.Equal(t, 6, l.first.next.value)
	require.Equal(t, 4, l.first.next.next.value)
	require.Equal(t, 8, l.first.next.next.next.value)
	require.Equal(t, 7, l.last.value)
	require.Equal(t, 5, Size(l))
}

func TestRemoveItem(t *testing.T) {
	l := New()

	PushBack(l, 1)
	PushBack(l, 2)
	PushBack(l, 3)

	require.Equal(t, false, RemoveItem(l, 4))

	require.Equal(t, true, RemoveItem(l, 1))
	require.Equal(t, true, RemoveItem(l, 3))

	require.Equal(t, 2, l.first.value)
	require.Equal(t, l.first, l.last)
}

func TestReverse(t *testing.T) {
	l := New()
	require.Equal(t, 0, Size(l))
	require.Nil(t, l.first)
	require.Nil(t, l.last)

	Reverse(l)

	require.Equal(t, 0, Size(l))
	require.Nil(t, l.first)
	require.Nil(t, l.last)

	PushBack(l, 1)
	Reverse(l)

	require.Equal(t, 1, Size(l))
	require.Equal(t, 1, l.first.value)
	require.Nil(t, l.first.next)
	require.Equal(t, 1, l.last.value)

	PushBack(l, 2)
	Reverse(l)

	require.Equal(t, 2, Size(l))
	require.Equal(t, 2, l.first.value)
	require.Equal(t, 1, l.first.next.value)
	require.Equal(t, 1, l.last.value)

	PushBack(l, 3)
	Reverse(l)

	require.Equal(t, 3, Size(l))
	require.Equal(t, 3, l.first.value)
	require.Equal(t, 1, l.first.next.value)
	require.Equal(t, 2, l.first.next.next.value)
	require.Equal(t, 2, l.last.value)
}
