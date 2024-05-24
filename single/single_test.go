package single

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var s []int
	for i := 1; i < 11; i++ {
		s = append(s, i)
	}
	l := New[int]()
	for _, v := range s {
		l.Append(v)
	}

	if l.head.value != s[0] {
		t.Fatalf("List head = %d; want %d\n", l.head.value, s[0])
	}

	if l.tail.value != s[len(s)-1] {
		t.Fatalf("List tail = %d; want %d\n", l.tail.value, s[len(s)-1])
	}

	if l.length != len(s) {
		t.Fatalf("List length = %d; want %d\n", l.length, len(s))
	}

	if l.pos == nil {
		t.Fatalf("List pos = nil; want not nil\n")
	}

	if l.pos != l.head {
		t.Fatalf("List pos = %d; want %d\n", l.pos.value, l.head.value)
	}

	if l.prev != nil {
		t.Fatalf("List prev != nil; want nil\n")
	}

	n := l.head
	c := 0
	for n != nil {
		c++
		n = n.next
	}
	if c != len(s) {
		t.Fatalf("List actual length = %d; want %d\n", c, len(s))
	}

	n = l.head
	for _, v := range s {
		if n.value != v {
			t.Fatalf("List node value = %d; want %d\n", n.value, v)
		}
		n = n.next
	}
}

func TestPrepend(t *testing.T) {
	var s []int
	for i := 1; i < 11; i++ {
		s = append(s, i)
	}
	l := New[int]()
	for _, v := range s {
		l.Prepend(v)
	}

	if l.head.value != s[len(s)-1] {
		t.Fatalf("List head = %d; want %d\n", l.head.value, s[len(s)-1])
	}

	if l.tail.value != s[0] {
		t.Fatalf("List tail = %d; want %d\n", l.tail.value, s[0])
	}

	if l.length != len(s) {
		t.Fatalf("List length = %d; want %d\n", l.length, len(s))
	}

	if l.pos == nil {
		t.Fatalf("List pos = nil; want not nil\n")
	}

	if l.pos != l.tail {
		t.Fatalf("List pos = %d; want %d\n", l.pos.value, l.tail.value)
	}

	if l.prev == nil {
		t.Fatalf("List prev = nil; want not nil\n")
	}

	if l.prev.value != s[1] {
		t.Fatalf("List prev = %d; want %d\n", l.prev.value, s[len(s)-2])
	}

	n := l.head
	c := 0
	for n != nil {
		c++
		n = n.next
	}
	if c != len(s) {
		t.Fatalf("List actual length = %d; want %d\n", c, len(s))
	}

	n = l.head
	for i := range s {
		v := s[len(s)-1-i]
		if n.value != v {
			t.Fatalf("List node value = %d; want %d\n", n.value, v)
		}
		n = n.next
	}
}

func TestLen(t *testing.T) {
	l := New[int]()
	if l.Len() != 0 {
		t.Fatalf("Len() = %d; want %d\n", l.Len(), 0)
	}
	for i := 0; i < 3; i++ {
		l.Append(i)
		if got, want := l.Len(), i+1; got != want {
			t.Fatalf("Len() = %d; want %d\n", got, want)
		}
	}
	if got, want := l.Len(), l.length; got != want {
		t.Fatalf("Len() = %d; want %d\n", got, want)
	}
}

func TestClear(t *testing.T) {
	l := New[int]()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Head()
	l.Next()
	l.Clear()
	if l.head != nil {
		t.Fatalf("List head != nil; want nil\n")
	}
	if l.tail != nil {
		t.Fatalf("List tail != nil; want nil\n")
	}
	if l.pos != nil {
		t.Fatalf("List pos != nil; want nil\n")
	}
	if l.prev != nil {
		t.Fatalf("List prev != nil; want nil\n")
	}
	if got, want := l.Len(), 0; got != want {
		t.Fatalf("Len() = %d; want %d\n", got, want)
	}
}

func TestOk(t *testing.T) {
	l := New[int]()
	l.Append(1)
	l.Append(2)
	l.Head()
	l.Next()
	if got, want := l.Ok(), true; got != want {
		t.Fatalf("Ok() = %t; want %t\n", got, want)
	}
	l.Next()
	if got, want := l.Ok(), false; got != want {
		t.Fatalf("Ok() = %t; want %t\n", got, want)
	}
}

func TestHead(t *testing.T) {
	l := New[int]()
	l.Prepend(1)
	l.Prepend(2)
	l.Head()
	if got, want := l.pos.value, l.head.value; got != want {
		t.Fatalf("List pos = %d; want = %d\n", got, want)
	}
	if l.prev != nil {
		t.Fatalf("List prev != nil; want nil\n")
	}
}

func TestNext(t *testing.T) {
	l := New[int]()
	var s []int
	for i := 1; i < 11; i++ {
		l.Append(i)
		s = append(s, i)
	}
	l.Head()
	for i := range s {
		if i == 0 {
			continue
		}
		l.Next()
		if got, want := l.prev.value, s[i-1]; got != want {
			t.Fatalf("List prev = %d; want = %d\n", got, want)
		}
		if got, want := l.pos.value, s[i]; got != want {
			t.Fatalf("List pos = %d; want = %d\n", got, want)
		}
	}
	l.Next()
	if got, want := l.prev.value, s[len(s)-1]; got != want {
		t.Fatalf("List prev = %d; want = %d\n", got, want)
	}
	if l.pos != nil {
		t.Fatalf("List pos != nil; want = nil\n")
	}
}

func TestValue(t *testing.T) {
	l := New[int]()
	l.Append(1)
	l.Head()
	if got, want := l.Value(), 1; got != want {
		t.Fatalf("Value() = %d; want = %d\n", got, want)
	}
}

func TestSetValue(t *testing.T) {
	l := New[int]()
	l.Append(1)
	l.Head()
	l.SetValue(2)
	if got, want := l.Value(), 2; got != want {
		t.Fatalf("Value() = %d; want = %d\n", got, want)
	}
}

func TestRemoveHead(t *testing.T) {
	l := New[int]()
	var s []int
	for i := 1; i < 11; i++ {
		s = append(s, i)
		l.Append(i)
	}

	for i := 0; i < len(s); i++ {
		l.RemoveHead()
		if got, want := l.Len(), len(s)-(i+1); got != want {
			t.Fatalf("Len() = %d; want %d\n", got, want)
		}

		n := l.head
		for _, v := range s[i+1:] {
			if n.value != v {
				t.Fatalf("List node value = %d; want %d\n", n.value, v)
			}
			n = n.next
		}
	}

	l.RemoveHead()
	if got, want := l.Len(), 0; got != want {
		t.Fatalf("Len() = %d; want %d\n", got, want)
	}

	l.Clear()
	for _, v := range s {
		l.Append(v)
	}
	l.Head()
	l.RemoveHead()
	if got, want := l.pos.value, 2; got != want {
		t.Fatalf("Link pos = %d; want %d\n", got, want)
	}

	l.Next()
	l.RemoveHead()
	if l.prev != nil {
		t.Fatalf("Link prev != nil; want nil\n")
	}
}

func TextRemove(t *testing.T) {
	l := New[int]()
	var s []int
	for i := 1; i < 11; i++ {
		s = append(s, i)
	}

	for _, v := range s {
		l.Append(v)
	}
	for i := 0; i < len(s)/2; i++ {
		l.Next()
		l.Remove()
		if got, want := l.Len(), len(s)-i-1; got != want {
			t.Fatalf("Len() = %d; want = %d\n", got, want)
		}
	}
	l.Head()
	for i := 0; i < len(s)/2; i++ {
		if got, want := l.Value(), s[i*2]; got != want {
			t.Fatalf("Value() = %d; want = %d\n", got, want)
		}
		l.Next()
	}

	l.Clear()
	l.Append(11)
	l.Remove()
	if got, want := l.Value(), 0; got != want {
		t.Fatalf("Value() = %d; want = %d\n", got, want)
	}
	if l.head != nil {
		t.Fatalf("List head != nil; want = nil\n")
	}
	if l.tail != nil {
		t.Fatalf("List tail != nil; want = nil\n")
	}
	if l.prev != nil {
		t.Fatalf("List prev != nil; want = nil\n")
	}

	l.Clear()
	l.Append(12)
	l.Append(13)
	l.Remove()
	if got, want := l.head.value, 13; got != want {
		t.Fatalf("List head = %d; want = %d\n", got, want)
	}
	if got, want := l.tail.value, 13; got != want {
		t.Fatalf("List tail = %d; want = %d\n", got, want)
	}
	if l.prev != nil {
		t.Fatalf("List prev != nil; want = nil\n")
	}

	l.Clear()
	l.Append(14)
	l.Append(15)
	l.Next()
	l.Remove()
	if got, want := l.head.value, 14; got != want {
		t.Fatalf("List head = %d; want = %d\n", got, want)
	}
	if got, want := l.tail.value, 14; got != want {
		t.Fatalf("List tail = %d; want = %d\n", got, want)
	}
	if got, want := l.prev.value, 14; got != want {
		t.Fatalf("List prev = %d; want = %d\n", got, want)
	}

	l.Clear()
	l.Append(16)
	l.Append(17)
	l.Append(18)
	l.Next()
	l.Remove()
	if got, want := l.head.value, 16; got != want {
		t.Fatalf("List head = %d; want = %d\n", got, want)
	}
	if got, want := l.tail.value, 18; got != want {
		t.Fatalf("List tail = %d; want = %d\n", got, want)
	}
	if got, want := l.prev.value, 16; got != want {
		t.Fatalf("List prev = %d; want = %d\n", got, want)
	}

	l.Clear()
	l.Remove()
	if got, want := l.Len(), 0; got != want {
		t.Fatalf("Len() = %d; want = %d\n", got, want)
	}
}
