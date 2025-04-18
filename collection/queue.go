package collection

import "errors"

// NOTE: A not so efficient queue implementation.
// A better way would be to use a "circular buffer" array/slice
// But at least incrementing by a fixed "large" amount is decent enough for this project
type Queue[T any] []*T

const GROW_INCREMENT = 100

func NewQueue[T any]() *Queue[T] {
	queue := Queue[T](make([]*T, 0, GROW_INCREMENT))
	return &queue
}

func (self *Queue[T]) Enqueue(el *T) {
	self.growIfNeeded()
	*self = append(*self, el)
}

func (self *Queue[T]) Dequeue() (*T, error) {
	if len(*self) == 0 {
		return nil, errors.New("Queue is empty")
	}

	el := (*self)[0]
	*self = (*self)[1:]
	return el, nil
}

func (self *Queue[T]) growIfNeeded() (grew bool) {
	queue := *self
	//This check works because when you Dequeue an element,
	//the capacity automatically shrinks
	if len(queue) == cap(queue) {
		newQueue := make([]*T, len(queue), cap(queue)+GROW_INCREMENT)
		copy(newQueue, queue)
		*self = newQueue
		return true
	}
	return false
}
