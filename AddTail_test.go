package genericlist

import "testing"

func TestLinkedList_AddTail(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		list := &LinkedList[int]{}

		// Add first element
		list.AddTail(1)
		if got := list.Count(); got != 1 {
			t.Fatalf("Count() = %d; want 1", got)
		}
		if list.head.data != 1 {
			t.Fatalf("head.data = %d; want 1", list.head.data)
		}
		if list.tail.data != 1 {
			t.Fatalf("tail.data = %d; want 1", list.tail.data)
		}
		if list.head.nextPtr != nil {
			t.Fatalf("head.nextPtr = %v; want nil", list.head.nextPtr)
		}
		if list.tail.prevPtr != nil {
			t.Fatalf("tail.prevPtr = %v; want nil", list.tail.prevPtr)
		}

		// Add second element
		list.AddTail(2)
		if got := list.Count(); got != 2 {
			t.Fatalf("Count() = %d; want 2", got)
		}
		if list.head.data != 1 {
			t.Fatalf("head.data = %d; want 1", list.head.data)
		}
		if list.head.nextPtr.data != 2 {
			t.Fatalf("head.nextPtr.data = %d; want 2", list.head.nextPtr.data)
		}
		if list.tail.data != 2 {
			t.Fatalf("tail.data = %d; want 2", list.tail.data)
		}
		if list.tail.prevPtr.data != 1 {
			t.Fatalf("tail.prevPtr.data = %d; want 1", list.tail.prevPtr.data)
		}
		if list.head.nextPtr.nextPtr != nil {
			t.Fatalf("head.nextPtr.nextPtr = %v; want nil", list.head.nextPtr.nextPtr)
		}

		// Add third element
		list.AddTail(3)
		if got := list.Count(); got != 3 {
			t.Fatalf("Count() = %d; want 3", got)
		}
		if list.head.data != 1 {
			t.Fatalf("head.data = %d; want 1", list.head.data)
		}
		if list.head.nextPtr.data != 2 {
			t.Fatalf("head.nextPtr.data = %d; want 2", list.head.nextPtr.data)
		}
		if list.head.nextPtr.nextPtr.data != 3 {
			t.Fatalf("head.nextPtr.nextPtr.data = %d; want 3", list.head.nextPtr.nextPtr.data)
		}
		if list.tail.data != 3 {
			t.Fatalf("tail.data = %d; want 3", list.tail.data)
		}
		if list.tail.prevPtr.data != 2 {
			t.Fatalf("tail.prevPtr.data = %d; want 2", list.tail.prevPtr.data)
		}
		if list.head.nextPtr.nextPtr.nextPtr != nil {
			t.Fatalf("head.nextPtr.nextPtr.nextPtr = %v; want nil", list.head.nextPtr.nextPtr.nextPtr)
		}
	})
	t.Run("make sure Add() resets sorted flag", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.sorted = true
		list.AddTail(2)
		if list.sorted {
			t.Fatal("sorted expects false")
		}
	})
}
