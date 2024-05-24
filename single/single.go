package single

type node[T any] struct {
	value T
	next  *node[T]
}

type List[T any] struct {
	length int
	head   *node[T]
	tail   *node[T]
	pos    *node[T]
	prev   *node[T]
}

func New[T any]() *List[T] {
	return new(List[T])
}

func (l *List[T]) Append(value T) {
	l.length++
	n := new(node[T])
	n.value = value
	if l.head == nil {
		l.head = n
		l.tail = n
		l.pos = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *List[T]) Prepend(value T) {
	l.length++
	n := new(node[T])
	n.value = value
	if l.head == nil {
		l.head = n
		l.tail = n
		l.pos = n
		return
	}
	if l.pos == l.head {
		l.prev = n
	}
	n.next = l.head
	l.head = n
}

func (l *List[T]) Len() int {
	return l.length
}

func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.pos = nil
	l.prev = nil
	l.length = 0
}

func (l *List[T]) Ok() bool {
	return l.pos != nil
}

func (l *List[T]) Head() {
	l.pos = l.head
	l.prev = nil
}

func (l *List[T]) Next() {
	if !l.Ok() {
		return
	}
	l.prev = l.pos
	l.pos = l.pos.next
}

func (l *List[T]) Value() (value T) {
	if l.Ok() {
		return l.pos.value
	}
	return value
}

func (l *List[T]) SetValue(value T) {
	if l.Ok() {
		l.pos.value = value
	}
}

func (l *List[T]) RemoveHead() {
	if l.length == 0 {
		return
	}
	if l.prev == l.head {
		l.prev = nil
	}
	if l.pos == l.head {
		l.pos = l.head.next
	}
	l.head = l.head.next
	if l.length == 1 {
		l.tail = nil
	}
	l.length--
}

func (l *List[T]) Remove() {
	if !l.Ok() {
		return
	}
	if l.head == l.pos {
		l.head = l.pos.next
	}
	if l.tail == l.pos {
		l.tail = l.prev
	}
	if l.prev != nil {
		l.prev.next = l.pos.next
	}
	l.pos = l.pos.next
	l.length--
}
