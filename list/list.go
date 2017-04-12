package list

type node struct {
	value int
	next  *node
}

type list struct {
	first *node
	size  int
	last  *node
}

func New() *list {
	return &list{nil, 0, nil}
}

func Size(l *list) int {
	return l.size
}

func Empty(l *list) bool {
	return l.size == 0
}

func Insert(l *list, index int, value int) bool {
	size := Size(l)
	newNode := &node{value, nil}

	if index < 0 {
		index = index + size
	}

	if index > size {
		return false
	}

	if index == 0 {
		prevFirst := l.first
		l.first = newNode
		if prevFirst != nil {
			l.first.next = prevFirst
		}

		if l.last == nil {
			l.last = l.first
		}
	} else {
		prev, _ := nodeAt(l, index-1)
		next, ok := nodeAt(l, index)

		prev.next = newNode
		if ok {
			newNode.next = next
		} else {
			l.last = newNode
		}
	}

	l.size++
	return true
}

func PushFront(l *list, value int) {
	Insert(l, 0, value)
}

func PushBack(l *list, value int) {
	Insert(l, Size(l), value)
}

func nodeAt(l *list, index int) (*node, bool) {
	size := Size(l)

	// allow negative index, means "from the end"
	if index < 0 {
		index = size + index - 1
	}

	if index >= size || index < 0 {
		return nil, false
	}

	// if it's the last element
	if index == size-1 {
		return l.last, true
	}

	cur := l.first
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur, true
}

func At(l *list, index int) (int, bool) {
	node, ok := nodeAt(l, index)
	if ok {
		return node.value, true
	} else {
		return 0, false
	}
}

func Remove(l *list, index int) bool {
	size := Size(l)

	if index >= size {
		return false
	}

	if size == 1 {
		l.first = nil
		l.last = nil
		l.size = 0
		return true
	}

	if index == 0 {
		l.first = l.first.next
	} else {
		prev, _ := nodeAt(l, index-1)
		prev.next = prev.next.next

		// update l.last when asked to remove the last element
		if index == size-1 {
			l.last = prev
		}
	}

	l.size--

	return true
}

func PopBack(l *list) (int, bool) {
	result, ok := Back(l)

	if ok {
		Remove(l, Size(l)-1)
	}

	return result, ok
}

func PopFront(l *list) (int, bool) {
	result, ok := Front(l)

	if ok {
		Remove(l, 0)
	}

	return result, ok
}

func RemoveItem(l *list, value int) bool {
	cur := l.first
	for i := 0; i < Size(l); i++ {
		if cur.value == value {
			return Remove(l, i)
		}
		cur = cur.next
	}

	return false
}

func Front(l *list) (int, bool) {
	if Size(l) == 0 {
		return 0, false
	} else {
		return l.first.value, true
	}
}

func Back(l *list) (int, bool) {
	if Size(l) == 0 {
		return 0, false
	} else {
		return l.last.value, true
	}
}

func Reverse(l *list) {
	first := l.first
	if first == nil {
		return
	}

	second := first.next
	if second == nil {
		return
	}

	prev := first
	cur := second
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	l.first, l.last = l.last, l.first
}
