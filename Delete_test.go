package genericlist

import "testing"

func TestLinkedList_Delete(t *testing.T) {
	list := &LinkedList[int]{}

	// Add elements to the list
	list.AddTail(1)
	list.AddTail(2)
	list.AddTail(3)

	// Initially, the count should be 3
	if got := list.count; got != 3 {
		t.Fatalf("Count() = %d; want 3", got)
	}

	// Delete the middle element (2)
	list.ToStart() // Move to the start
	if list.curr.data != 1 {
		t.Fatalf("curr.data = %d; want 1", list.curr.data)
	}
	list.Next() // Move to the second element
	if list.curr.data != 2 {
		t.Fatalf("curr.data = %d; want 2", list.curr.data)
	}
	list.Delete()

	// After deleting 2, count should be 2
	if got := list.count; got != 2 {
		t.Fatalf("Count() = %d; want 2", got)
	}
	if list.head.data != 1 {
		t.Fatalf("head.data = %d; want 1", list.head.data)
	}
	if list.tail.data != 3 {
		t.Fatalf("tail.data = %d; want 3", list.tail.data)
	}
	if list.head.nextPtr.data != 3 {
		t.Fatalf("head.nextPtr.data = %d; want 3", list.head.nextPtr.data)
	}
	if list.tail.prevPtr.data != 1 {
		t.Fatalf("tail.prevPtr.data = %d; want 1", list.tail.prevPtr.data)
	}

	// Delete the head element (1)
	list.ToStart() // Move to the start
	list.Delete()

	// After deleting 1, count should be 1
	if got := list.count; got != 1 {
		t.Fatalf("Count() = %d; want 1", got)
	}
	if list.head.data != 3 {
		t.Fatalf("head.data = %d; want 3", list.head.data)
	}
	if list.tail.data != 3 {
		t.Fatalf("tail.data = %d; want 3", list.tail.data)
	}
	if list.head.nextPtr != nil {
		t.Fatalf("head.nextPtr = %v; want nil", list.head.nextPtr)
	}

	// Delete the last element (3)
	list.Delete()

	// After deleting 3, count should be 0
	if got := list.count; got != 0 {
		t.Fatalf("Count() = %d; want 0", got)
	}
	if list.head != nil {
		t.Fatalf("head = %v; want nil", list.head)
	}
	if list.tail != nil {
		t.Fatalf("tail = %v; want nil", list.tail)
	}
	if list.curr != nil {
		t.Fatalf("curr = %v; want nil", list.curr)
	}
}
