package collection

import (
	"testing"
)

func TestQueue(t *testing.T) {	
	queue := NewQueue[int]()

	queue.Enqueue(1)
	queue.Enqueue(2)

	one, _ := queue.Dequeue()
	if one != 1 {
		t.Error("Dequeue not returning correct element")
	}

	queue.Enqueue(3)

	two, _ := queue.Dequeue()
	if two != 2 {
		t.Error("Dequeue not returning correct element")
	}
	three, _ := queue.Dequeue()
	if three != 3 {
		t.Error("Dequeue not returning correct element")
	}

	_, err := queue.Dequeue()
	if err == nil {
		t.Error("Queue is empty when attempting Dequeue. Should have returned an error.")
	}
}
