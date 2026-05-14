package list

import "iter"

type Node[T any] struct {
	v    T
	next *Node[T]
}

type Head[T any] struct {
	n   *Node[T]
	cmp func(T, T) bool
	len int
}

func (h *Head[T]) Append(v T) {
	lappend(&h.n, v)
	h.len++
}

func (h *Head[T]) Prepend(v T) {
	prepend(&h.n, v)
	h.len++
}

func all[T any](n *Node[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for ; n != nil; n = n.next {
			if !yield(n.v) {
				return
			}
		}
	}
}

func prepend[T any](n **Node[T], v T) {
	*n = &Node[T]{v: v, next: *n}
}

func lappend[T any](n **Node[T], v T) {
	for ; *n != nil; n = &(*n).next {
	}
	*n = &Node[T]{v: v}
}

func ldelete[T any](n **Node[T], v T, cmp func(T, T) bool) {
	for ; *n != nil; n = &(*n).next {
		if !cmp((*n).v, v) {
			continue
		}
		*n = (*n).next
		return
	}
}
