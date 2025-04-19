package collection

import "errors"

// NOTE: A not so efficient queue implementation.
// A better way would be to use a "circular buffer" array/slice
// But at least incrementing by a fixed "large" amount is decent enough for this project
type Queue[T any] []T

const GROW_INCREMENT = 1024

func NewQueue[T any]() *Queue[T] {
	queue := Queue[T](make([]T, 0, GROW_INCREMENT))
	return &queue
}

func (self *Queue[T]) Enqueue(el T) {
	self.growIfNeeded()
	*self = append(*self, el)
}

func (self *Queue[T]) Dequeue() (T, error) {
	if self.IsEmpty() {
		return *new(T), errors.New("Queue is empty. Always check IsEmpty() before attempting to Dequeue.")
	}

	el := (*self)[0]
	*self = (*self)[1:]
	return el, nil
}

func (self *Queue[T]) IsEmpty() bool {
	return len(*self) == 0
}

func (self *Queue[T]) growIfNeeded() {
	queue := *self
	//This check works because when you Dequeue an element,
	//the capacity automatically shrinks
	if len(queue) == cap(queue) {
		newQueue := make([]T, len(queue), cap(queue)+GROW_INCREMENT)
		copy(newQueue, queue)
		*self = newQueue
	}
}
