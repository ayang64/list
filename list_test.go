package list

import "testing"

func TestAll(t *testing.T) {
	var l *Node[int]

	lappend(&l, 1)
	lappend(&l, 2)
	lappend(&l, 12)
	lappend(&l, 42)
	lappend(&l, -8)
	lappend(&l, -128)
	lappend(&l, 13)

	for v := range all(l) {
		t.Logf("%#v", v)
	}
}

func TestPrepend(t *testing.T) {
	var l *Node[int]
	prepend(&l, 1)
	prepend(&l, 2)
	prepend(&l, 12)
	prepend(&l, 42)
	prepend(&l, -8)
	prepend(&l, -128)
	prepend(&l, 13)
	for ; l != nil; l = l.next {
		t.Logf("%#v", l.v)
	}
}

func TestLappend(t *testing.T) {
	var l *Node[int]

	lappend(&l, 1)
	lappend(&l, 2)
	lappend(&l, 12)
	lappend(&l, 42)
	lappend(&l, -8)
	lappend(&l, -128)
	lappend(&l, 13)

	for ; l != nil; l = l.next {
		t.Logf("%#v", l.v)
	}
}

func TestLdelete(t *testing.T) {
	var l *Node[int]

	lappend(&l, 1)
	lappend(&l, 2)
	lappend(&l, 12)
	lappend(&l, 42)
	lappend(&l, -8)
	lappend(&l, -128)
	lappend(&l, 13)

	ldelete(&l, 42, func(a, b int) bool {
		return a == b
	})

	ldelete(&l, 1, func(a, b int) bool {
		return a == b
	})

	ldelete(&l, 13, func(a, b int) bool {
		return a == b
	})

	for ; l != nil; l = l.next {
		t.Logf("%#v", l.v)
	}
}
