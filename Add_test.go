package genericlist

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		// Create a new linked list
		list := &LinkedList[int]{}

		if list.sorted {
			t.Fatal("sorted expects false")
		}

		// Add some nodes to the list
		for i := 1; i <= 10; i++ {
			list.Add(i)
			if list.sorted {
				t.Fatal("sorted expects false")
			}
		}

		if list.count != 10 {
			t.Fatal("count mismatch")
		}

		if list.head.data != 1 {
			t.Fatal("head.data mismatch 1")
		}
		if list.tail.data != 10 {
			t.Fatal("tail.data mismatch 3")
		}
		if list.curr.data != 10 {
			t.Fatal("curr.data mismatch 3")
		}
		if list.head.nextPtr.data != 2 {
			t.Fatal("head.nextPtr.data mismatch 2")
		}
		if list.head.nextPtr.nextPtr.data != 3 {
			t.Fatal("head.nextPtr.nextPtr.data mismatch 3")
		}
		if list.head.nextPtr.prevPtr.data != 1 {
			t.Fatal("head.nextPtr.prevPtr.data mismatch 1")
		}
	})
	t.Run("make sure Add() resets sorted flag", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.sorted = true
		list.Add(2)
		if list.sorted {
			t.Fatal("sorted expects false")
		}
	})
}
