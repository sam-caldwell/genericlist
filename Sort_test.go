package genericlist

import "testing"

func TestLinkedList_Sort(t *testing.T) {
	// Helper function for less comparison in ascending order
	AscendingOrder := func(a, b int) bool {
		return a > b
	}

	// Helper function for less comparison in descending order
	DescendingOrder := func(a, b int) bool {
		return a < b
	}

	// Helper function to get the list data as a slice
	getListData := func(list *LinkedList[int]) []int {
		var data []int
		list.curr = list.head
		for list.curr != nil {
			data = append(data, list.Get())
			if !list.moveNext() {
				break
			}
		}
		return data
	}

	// Test sorting in ascending order
	t.Run("Sort Ascending", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add elements in a random order
		list.AddTail(3)
		list.AddTail(1)
		list.AddTail(2)

		// Sort in ascending order
		list.Sort(AscendingOrder)

		expected := []int{1, 2, 3}
		got := getListData(list)

		for i, v := range expected {
			if got[i] != v {
				t.Errorf("Sort Ascending: got %v; want %v", got, expected)
				return
			}
		}
	})

	// Test sorting in descending order
	t.Run("Sort Descending", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add elements in a random order
		list.AddTail(1)
		list.AddTail(3)
		list.AddTail(2)

		// Sort in descending order
		list.Sort(DescendingOrder)

		expected := []int{3, 2, 1}
		got := getListData(list)

		for i, v := range expected {
			if got[i] != v {
				t.Errorf("Sort Descending: got %v; want %v", got, expected)
				return
			}
		}
	})
}
