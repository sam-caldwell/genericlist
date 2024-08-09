package genericlist

import "testing"

func TestLinkedList_ToList(t *testing.T) {
	// Helper function to compare slices
	compareSlices := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	// Test with an empty list
	t.Run("Empty List", func(t *testing.T) {
		list := &LinkedList[int]{}

		got := list.ToList()
		expected := []int{}

		if !compareSlices(got, expected) {
			t.Errorf("Empty List: got %v; want %v", got, expected)
		}
	})

	// Test with a single element
	t.Run("Single Element", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.AddTail(1)

		got := list.ToList()
		expected := []int{1}

		if !compareSlices(got, expected) {
			t.Errorf("Single Element: got %v; want %v", got, expected)
		}
	})

	// Test with multiple elements
	t.Run("Multiple Elements", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.AddTail(1)
		list.AddTail(2)
		list.AddTail(3)

		got := list.ToList()
		expected := []int{1, 2, 3}

		if !compareSlices(got, expected) {
			t.Errorf("Multiple Elements: got %v; want %v", got, expected)
		}
	})

	// Test with list having unsorted elements
	t.Run("Unsorted Elements", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.AddTail(5)
		list.AddTail(3)
		list.AddTail(4)
		list.AddTail(1)

		got := list.ToList()
		expected := []int{5, 3, 4, 1}

		if !compareSlices(got, expected) {
			t.Errorf("Unsorted Elements: got %v; want %v", got, expected)
		}
	})
}
