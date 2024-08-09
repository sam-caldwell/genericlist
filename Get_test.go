package genericlist

import "testing"

func TestLinkedList_Get(t *testing.T) {
	// Test when list is empty
	t.Run("Empty List", func(t *testing.T) {
		list := &LinkedList[int]{}
		value := list.Get()
		if value != 0 {
			t.Fatal("expect zero value")
		}
	})

	// Test when list has elements and current node is set
	t.Run("Non-Empty List", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add elements and set the current node
		list.AddTail(1)
		list.AddTail(2)
		list.AddTail(3)

		// Move to the second node
		list.curr = list.head
		if got := list.Get(); got != 1 {
			t.Fatalf("Get() = %d; want 1", got)
		}
		list.curr = list.curr.nextPtr
		if got := list.Get(); got != 2 {
			t.Fatalf("Get() = %d; want 2", got)
		}
		list.curr = list.curr.nextPtr
		if got := list.Get(); got != 3 {
			t.Fatalf("Get() = %d; want 3", got)
		}
	})

	// Test when list has elements but current node is nil
	t.Run("Current Node Nil", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.curr = nil
		list.Get()
	})
}
