package genericlist

import "testing"

func TestLinkedList_BinarySearch(t *testing.T) {
	// Helper function to set up a linked list with integers
	setupList := func(values []int) *LinkedList[int] {
		list := &LinkedList[int]{}
		for _, v := range values {
			list.AddTail(v)
		}
		return list
	}

	t.Run("Found at start", func(t *testing.T) {
		list := setupList([]int{1, 2, 3, 4, 5})
		found := list.BinarySearch(1)
		if !found {
			t.Errorf("Expected to find 1, but did not.")
		}
	})

	t.Run("Found in middle", func(t *testing.T) {
		list := setupList([]int{1, 2, 3, 4, 5})
		found := list.BinarySearch(3)
		if !found {
			t.Errorf("Expected to find 3, but did not.")
		}
	})

	t.Run("Found at end", func(t *testing.T) {
		list := setupList([]int{1, 2, 3, 4, 5})
		found := list.BinarySearch(5)
		if !found {
			t.Errorf("Expected to find 5, but did not.")
		}
	})

	t.Run("Not found", func(t *testing.T) {
		list := setupList([]int{1, 2, 3, 4, 5})
		found := list.BinarySearch(6)
		if found {
			t.Errorf("Expected not to find 6, but did.")
		}
	})

	t.Run("Empty list", func(t *testing.T) {
		list := &LinkedList[int]{}
		found := list.BinarySearch(1)
		if found {
			t.Errorf("Expected not to find 1 in empty list, but did.")
		}
	})

	t.Run("Single element - found", func(t *testing.T) {
		list := setupList([]int{1})
		found := list.BinarySearch(1)
		if !found {
			t.Errorf("Expected to find 1 in single element list, but did not.")
		}
	})

	t.Run("Single element - not found", func(t *testing.T) {
		list := setupList([]int{1})
		found := list.BinarySearch(2)
		if found {
			t.Errorf("Expected not to find 2 in single element list, but did.")
		}
	})

	t.Run("Unsorted list - found after sorting", func(t *testing.T) {
		list := setupList([]int{5, 1, 3, 4, 2})
		found := list.BinarySearch(3)
		if !found {
			t.Errorf("Expected to find 3 after sorting, but did not.")
		}
	})

	t.Run("Unsorted list - not found after sorting", func(t *testing.T) {
		list := setupList([]int{5, 1, 3, 4, 2})
		found := list.BinarySearch(6)
		if found {
			t.Errorf("Expected not to find 6 after sorting, but did.")
		}
	})
}
