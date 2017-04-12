package arrayInt

import (
	"fmt"
	"math"
)

type any interface{}

type array struct {
	array []any
	size  int
	cap   int
}

func Create(initialCap int) *array {
	// Covert initial capacity into power of 2. Starting from 16
	cap := 16
	for cap < initialCap {
		cap *= 2
	}

	return &array{make([]any, 0, cap), 0, cap}
}

func Cap(arr *array) int {
	return arr.cap
}

func Size(arr *array) int {
	return arr.size
}

func IsEmpty(arr *array) bool {
	return arr.size == 0
}

func At(arr *array, index int) any {
	size := Size(arr)
	if index >= size {
		panic(fmt.Sprintf(
			"Index out of bound. The size of the array was %d, but the requested index was %d",
			size,
			index,
		))
	}

	return arr.array[index]
}

func Set(arr *array, index int, item any) {
	size := Size(arr)
	if index >= size {
		panic(fmt.Sprintf(
			"Index out of bound. The size of the array was %d, but the requested index was %d",
			size,
			index,
		))
	}

	arr.array[index] = item
}

func resize(arr *array, newCapacity int) {
	if newCapacity == Cap(arr) {
		return
	}

	size := Size(arr)

	if newCapacity < size {
		panic(fmt.Sprintf(
			"Tried to resize an array with size %d to the capacity %d which is smaller",
			size,
			newCapacity,
		))
	}

	newArray := make([]any, size, newCapacity)

	// copy elements from the old array to the new one
	for i := 0; i < size; i++ {
		newArray[i] = At(arr, i)
	}

	(*arr).array = newArray
	(*arr).cap = newCapacity

	return
}

func Push(arr *array, item any) {
	cap := Cap(arr)
	size := Size(arr)

	// we are at full capacity
	if cap == size {
		resize(arr, cap*2)
	}
	arr.array = arr.array[:size+1]
	arr.size = size + 1

	arr.array[size] = item
}

func Insert(arr *array, index int, item any) {
	cap := Cap(arr)
	size := Size(arr)
	sizeAbs := int(math.Abs(float64(size)))

	if index > sizeAbs {
		panic(fmt.Sprintf(
			"Index out of bound. The size of the array was %d, but the requested index was %d",
			size,
			index,
		))
	}

	// allow negative index, means "from the end"
	if index < 0 {
		index = size + index
	}

	// we are at full capacity
	if cap == size {
		resize(arr, cap*2)
	}
	arr.array = arr.array[:size+1]
	arr.size = size + 1

	for i := Size(arr) - 2; i >= index; i-- {
		arr.array[i+1] = arr.array[i]
	}

	arr.array[index] = item
}

func Prepend(arr *array, item any) {
	Insert(arr, 0, item)
}

func Pop(arr *array) (result any) {
	size := Size(arr)
	index := size - 1 // last element
	if size <= 0 {
		panic("Tried to call Pop() on an empty array.")
	}

	result = arr.array[index]

	Delete(arr, index)

	return
}

func Delete(arr *array, index int) {
	size := Size(arr)
	cap := Cap(arr)

	if index >= size {
		panic(fmt.Sprintf(
			"Index out of bound. The size of the array was %d, but the requested index was %d",
			size,
			index,
		))
	}

	for i := index; i < size-1; i++ {
		arr.array[i] = arr.array[i+1]
	}

	arr.array = arr.array[:size-1]
	arr.size = size - 1

	if Size(arr)*4 <= cap && cap/2 >= 16 {
		resize(arr, cap/2)
	}
}

func Find(arr *array, item any) (int, bool) {
	for i := 0; i < Size(arr); i++ {
		if At(arr, i) == item {
			return i, true
		}
	}

	return 0, false
}

func Remove(arr *array, item any) bool {
	index, ok := Find(arr, item)

	if ok {
		Delete(arr, index)
	}

	return ok
}
