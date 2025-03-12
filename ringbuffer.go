package goex

type RingBuffer[T any] struct {
	data []T
	size int
	w_i  int
}

func NewRingBuffer[T any](size int) RingBuffer[T] {
	return RingBuffer[T]{
		data: make([]T, size),
		size: 0,
		w_i:  0,
	}
}

func (b RingBuffer[T]) Len() int {
	return b.size
}

func (b RingBuffer[T]) Last() (int, T) {
	i := (b.size + len(b.data) - 1) % len(b.data)
	return i, b.Get(i)
}

func (b RingBuffer[T]) Get(i int) T {
	return b.data[i%len(b.data)]
}

func (b RingBuffer[T]) Set(i int, v T) {
	b.data[i%len(b.data)] = v
}

func (b RingBuffer[T]) Put(v T) RingBuffer[T] {
	b.data[b.w_i] = v
	return RingBuffer[T]{
		data: b.data,
		size: min(len(b.data), b.size+1),
		w_i:  (b.w_i + 1) % len(b.data),
	}
}
