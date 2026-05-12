package async

import (
	"iter"
	"slices"
)

type List[T any] struct {
	Items     iter.Seq[T]
	notifiers []*ListNotifier[T]
}

type ListNotifier[T any] struct {
	list      *List[T]
	OnAdded   func(v T)
	OnRemoved func(v T)
	OnUpdated func()
}

func (l *List[T]) Subscribe() *ListNotifier[T] {
	ln := &ListNotifier[T]{list: l}
	l.notifiers = append(l.notifiers, ln)
	return ln
}

func (l *List[T]) Unsubscribe(ln *ListNotifier[T]) {
	l.notifiers = slices.DeleteFunc(l.notifiers, func(n *ListNotifier[T]) bool {
		return n == ln
	})
}

func (l *List[T]) Add(v T) {
	for _, n := range l.notifiers {
		if n.OnAdded != nil {
			n.OnAdded(v)
		}
		if n.OnUpdated != nil {
			n.OnUpdated()
		}
	}
}

func (l *List[T]) Remove(v T) {
	for _, n := range l.notifiers {
		if n.OnRemoved != nil {
			n.OnRemoved(v)
		}
		if n.OnUpdated != nil {
			n.OnUpdated()
		}
	}
}

func (l *ListNotifier[T]) Defer() {
	l.list.Unsubscribe(l)
}

func (l *ListNotifier[T]) Unsubscribe() {
	l.list.Unsubscribe(l)
}
