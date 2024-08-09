package genericlist

import "testing"

func TestLinkedList_Count(t *testing.T) {

	t.Run("test empty", func(t *testing.T) {
		list := &LinkedList[int]{}
		// Initially, the count should be 0
		if got := list.Count(); got != 0 {
			t.Fatalf("Count() = %d; want 0", got)
		}
	})

	t.Run("test growing list", func(t *testing.T) {
		list := &LinkedList[int]{}
		for i := 1; i <= 10; i++ {
			list.AddTail(i)
			if list.count != list.Count() {
				t.Fatalf("List[%d] = %d; want %d", i, list.count, list.Count())
			}
			if list.count != uint(i) {
				t.Fatalf("List[%d] = %d; want %d", i, list.count, i)
			}
		}
	})

	t.Run("test with deletes", func(t *testing.T) {
		list := &LinkedList[int]{}
		for i := 1; i <= 10; i++ {
			list.AddTail(i)
			if list.count != list.Count() {
				t.Fatalf("List[%d] = %d; want %d", i, list.count, list.Count())
			}
			if list.count != uint(i) {
				t.Fatalf("List[%d] = %d; want %d", i, list.count, i)
			}
		}
		// Remove an element
		list.Delete() // Assuming Delete() removes the current node; in this case, it will remove the node with data 10

		// After removing 1 element, count should be 9
		if got := list.Count(); got != 9 {
			t.Fatalf("Count() = %d; want 9", got)
		}

	})

}
