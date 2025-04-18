package collection

import (
	"testing"
)

func TestQueue(t *testing.T) {	
	queue := NewQueue[int]()

	val1 := 1
	val2 := 2
	queue.Enqueue(&val1)
	queue.Enqueue(&val2)

	one, _ := queue.Dequeue()
	if *one != val1 {
		t.Error("Dequeue not returning correct element")
	}

	val3 := 3
	queue.Enqueue(&val3)

	two, _ := queue.Dequeue()
	if *two != val2 {
		t.Error("Dequeue not returning correct element")
	}
	three, _ := queue.Dequeue()
	if *three != val3 {
		t.Error("Dequeue not returning correct element")
	}

	_, err := queue.Dequeue()
	if err == nil {
		t.Error("Queue is empty when attempting Dequeue. Should have returned an error.")
	}
}
