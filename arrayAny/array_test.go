package arrayInt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	// keys are "initialCap" and values are expected real capacity
	testMap := map[int]int{
		0:   16,
		10:  16,
		16:  16,
		17:  32,
		100: 128,
	}

	for initialCap, expectedCap := range testMap {
		gotCap := Cap(Create(initialCap))
		if gotCap != expectedCap {
			t.Errorf(
				"Expected real capacity %d from initial capacity %d, got %d",
				expectedCap,
				initialCap,
				gotCap,
			)
		}
	}
}

func TestCap(t *testing.T) {
	arr := Create(16)
	expected := arr.cap
	actual := Cap(arr)
	require.Equal(t, expected, actual, "Cap(arr) should return the actual capacity of the array")
}

func TestSize(t *testing.T) {
	arr := Create(16)
	expected := arr.size
	actual := Size(arr)
	require.Equal(t, expected, actual, "Size(arr) should return the actual size of the array")
}

func TestIsEmpty(t *testing.T) {
	arr := Create(16)
	require.Equal(t, 0, Size(arr), "The size of initially created array should be 0")

	require.Equal(t, true, IsEmpty(arr), "IsEmpty should return true when the size of the array is 0")

	// HACK: it's actually impossible to change the size of the array this way, only by appending
	arr.size = 10

	require.Equal(t, false, IsEmpty(arr), "IsEmpty should return false when the size of the array is greater than 0")
}

func TestAtPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("At() should panic when out of bounds, but it didn't")
		}
	}()

	arr := Create(16)
	At(arr, 0)
}

func TestAt(t *testing.T) {
	arr := Create(16)

	expected := 4
	arr.array = append(arr.array, expected)
	arr.size = 1
	actual := At(arr, 0)

	require.Equal(t, expected, actual)
}

func TestSetPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Set() should panic when out of bounds, but it didn't")
		}
	}()

	arr := Create(16)
	Set(arr, 0, 1)
}

func TestSet(t *testing.T) {
	arr := Create(16)

	Insert(arr, 0, 4)
	require.Equal(t, 4, At(arr, 0))

	Set(arr, 0, 5)
	require.Equal(t, 5, At(arr, 0))
}

func TestResizePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Resize() should panic when the new capacity is smaller than the size, but it didn't")
		}
	}()

	oldCapacity := 16
	newCapacity := 15

	arr := Create(oldCapacity)

	// emulate filling all available space
	arr.size = oldCapacity

	resize(arr, newCapacity)
}

func TestResize(t *testing.T) {
	oldCapacity := 16
	newCapacity := 33

	arr := Create(oldCapacity)

	require.Equal(t, oldCapacity, cap(arr.array))
	require.Equal(t, oldCapacity, Cap(arr))

	resize(arr, newCapacity)

	require.Equal(t, newCapacity, cap(arr.array))
	require.Equal(t, newCapacity, Cap(arr))
}

func TestPush(t *testing.T) {
	arr := Create(16)

	// capacity doesn't change until there is free space for new elements
	for i := 0; i < 16; i++ {
		Push(arr, 1)
		require.Equal(t, 16, Cap(arr))
		require.Equal(t, i+1, Size(arr))
	}

	// capacity changes when trying to push a new element when the array is full
	require.Equal(t, 16, Cap(arr))
	require.Equal(t, 16, Size(arr))
	Push(arr, 1)
	require.Equal(t, 32, Cap(arr))
	require.Equal(t, 17, Size(arr))
}

func TestInsertPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Insert() should panic the index is out of bounds")
		}
	}()

	arr := Create(16)
	Insert(arr, 16, 42)
}

func TestInsert(t *testing.T) {
	arr := Create(16)

	// Fill a half of the array in a simple "push" fashion
	for i := 0; i < 8; i++ {
		Insert(arr, i, 1)
	}

	require.Equal(t, 16, Cap(arr))
	require.Equal(t, 8, Size(arr))

	require.Equal(t, []int{1, 1, 1, 1, 1, 1, 1, 1}, toInt(arr.array))

	// Insert 2 just in the middle of the array, expecting the elements to shift
	Insert(arr, 3, 2)
	require.Equal(t, 9, Size(arr))
	require.Equal(t, []int{1, 1, 1, 2, 1, 1, 1, 1, 1}, toInt(arr.array))

	// Test negative index
	Insert(arr, -3, 2)
	require.Equal(t, 10, Size(arr))
	require.Equal(t, []int{1, 1, 1, 2, 1, 1, 2, 1, 1, 1}, toInt(arr.array))

	// Make the size of the array > 16
	for i := 10; i < 18; i++ {
		Insert(arr, i, 3)
	}

	require.Equal(t, 32, Cap(arr))
	require.Equal(t, 18, Size(arr))
}

func TestPrepend(t *testing.T) {
	arr := Create(16)
	require.Equal(t, 16, Cap(arr))
	require.Equal(t, 0, Size(arr))

	Prepend(arr, 1)
	require.Equal(t, 1, Size(arr))
	require.Equal(t, []int{1}, toInt(arr.array))

	Prepend(arr, 2)
	require.Equal(t, 2, Size(arr))
	require.Equal(t, []int{2, 1}, toInt(arr.array))
}

func TestPopPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop() should panic when the array is empty")
		}
	}()

	arr := Create(16)
	require.Equal(t, 0, Size(arr))
	Pop(arr)
}

func TestPop(t *testing.T) {
	arr := Create(16)
	Push(arr, 1)
	Push(arr, 2)
	Push(arr, 3)

	require.Equal(t, 3, Pop(arr))
	require.Equal(t, 2, Pop(arr))
	require.Equal(t, 1, Pop(arr))
}

func TestDeletePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Delete() should panic when the is out of bounds")
		}
	}()

	arr := Create(16)
	Push(arr, 1)

	Delete(arr, 1)
}

func TestDelete(t *testing.T) {
	arr := Create(16)
	Push(arr, 1)
	Push(arr, 2)
	Push(arr, 3)
	require.Equal(t, 3, Size(arr))

	Delete(arr, 1)
	require.Equal(t, []int{1, 3}, toInt(arr.array))
	require.Equal(t, 2, Size(arr))

	Delete(arr, 0)
	require.Equal(t, []int{3}, toInt(arr.array))
	require.Equal(t, 1, Size(arr))

	Delete(arr, 0)
	require.Equal(t, []int{}, toInt(arr.array))
	require.Equal(t, 0, Size(arr))
}

func TestDeleteCapacity(t *testing.T) {
	arr := Create(32)

	// capacity doesn't change until there is free space for new elements
	for i := 0; i < 32; i++ {
		Push(arr, 1)
		require.Equal(t, 32, Cap(arr))
		require.Equal(t, i+1, Size(arr))
	}

	for i := 0; i < 23; i++ {
		Delete(arr, 0)
		require.Equal(t, 32, Cap(arr))
		require.Equal(t, 32-i-1, Size(arr))
	}

	require.Equal(t, 9, Size(arr))
	require.Equal(t, 32, Cap(arr))

	Delete(arr, 0)

	require.Equal(t, 8, Size(arr))
	require.Equal(t, 16, Cap(arr))
}

func TestFind(t *testing.T) {
	arr := Create(16)
	Push(arr, 1)
	Push(arr, 2)
	Push(arr, 3)

	var index int
	var ok bool

	index, ok = Find(arr, 4)
	require.Equal(t, false, ok)
	require.Equal(t, 0, index)

	index, ok = Find(arr, 2)
	require.Equal(t, true, ok)
	require.Equal(t, 1, index)
}

func TestRemove(t *testing.T) {
	arr := Create(16)
	Push(arr, 1)
	Push(arr, 2)
	Push(arr, 3)
	require.Equal(t, 3, Size(arr))

	var ok bool

	ok = Remove(arr, 4)
	require.Equal(t, false, ok)
	require.Equal(t, []int{1, 2, 3}, toInt(arr.array))
	require.Equal(t, 3, Size(arr))

	ok = Remove(arr, 2)
	require.Equal(t, true, ok)
	require.Equal(t, []int{1, 3}, toInt(arr.array))
	require.Equal(t, 2, Size(arr))
}

func TestDifferentTypes(t *testing.T) {
	arr := Create(16)
	Push(arr, 1)
	Push(arr, "2")
	Push(arr, false)

	require.Equal(t, 1, At(arr, 0).(int))
	require.Equal(t, "2", At(arr, 1).(string))
	require.Equal(t, false, At(arr, 2).(bool))

	_, ok := At(arr, 0).(string)

	require.Equal(t, false, ok)
}

func toInt(input []any) []int {
	arrInt := make([]int, 0)
	for _, v := range input {
		arrInt = append(arrInt, v.(int))
	}
	return arrInt
}
