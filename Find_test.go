package genericlist

import (
	"math/rand/v2"
	"testing"
)

// Helper function for comparison
func intCompare(lhs, rhs int) bool {
	return lhs == rhs
}

func TestLinkedList_Find(t *testing.T) {
	// Scenario 1: Sorted List
	t.Run("Sorted List", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add elements to the list (sorted)
		for i := 0; i < 100; i++ {
			list.Add(i)
		}

		compare := intCompare

		// Test finding an existing item
		if found := list.Find(50, compare); !found {
			t.Fatalf("Find(50) = %v; want true", found)
		}

		// Test finding a non-existing item
		if found := list.Find(-1, compare); found {
			t.Fatalf("Find(-1) = %v; want false", found)
		}
		if found := list.Find(200, compare); found {
			t.Fatalf("Find(200) = %v; want false", found)
		}
	})

	// Scenario 2: Random Unsorted List
	t.Run("Unsorted List", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add elements to the list (unsorted)
		existingNumber := 0
		for i := 0; i < 100; i++ {
			n := rand.Int()
			list.Add(n)
			if i == 50 {
				existingNumber = n
			}
		}

		compare := intCompare

		// Test finding an existing item
		if found := list.Find(existingNumber, compare); !found {
			t.Fatalf("Find(%d) = %v; want true", existingNumber, found)
		}

		// Test finding a non-existing item
		if found := list.Find(-4, compare); found {
			t.Fatalf("Find(-4) = %v; want false", found)
		}

		// Test finding the first item after traversal
		if found := list.Find(list.head.nextPtr.data, compare); !found {
			t.Fatalf("Find(%d) = %v; want true", list.head.nextPtr.data, found)
		}

		// Test finding the second item after traversal
		list.curr = list.head.nextPtr
		list.moveNext() // Move to the second element
		if found := list.Find(list.head.data, compare); !found {
			t.Fatalf("Find(%d) = %v; want true", list.head.data, found)
		}
	})
}
