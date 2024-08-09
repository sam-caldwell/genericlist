package genericlist

import "testing"

func TestLinkedList_AddHead(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		list := &LinkedList[int]{}

		list.AddHead(1)
		if got := list.Count(); got != 1 {
			t.Fatalf("Count() = %d; want 1", got)
		}
		if list.head.prevPtr != nil {
			t.Fatalf("list.head.prevPtr = %v; want nil", list.head.prevPtr)
		}
		if list.head.data != 1 {
			t.Fatal("expected head 1")
		}
		if list.tail.data != 1 {
			t.Fatal("expected tail 1")
		}

		list.AddHead(2)
		if got := list.Count(); got != 2 {
			t.Fatalf("Count() = %d; want 2", got)
		}
		if list.head.prevPtr != nil {
			t.Fatalf("list.head.prevPtr = %v; want nil", list.head.prevPtr)
		}
		if list.head.data != 2 {
			t.Fatal("expected head 2")
		}
		if list.tail.data != 1 {
			t.Fatal("expected tail 1")
		}

		list.AddHead(3)
		if got := list.Count(); got != 3 {
			t.Fatalf("Count() = %d; want 3", got)
		}
		if list.head.prevPtr != nil {
			t.Fatalf("list.head.prevPtr = %v; want nil", list.head.prevPtr)
		}
		if list.head.data != 3 {
			t.Fatal("expected 3")
		}
		if list.head.data != 3 {
			t.Fatal("expected head 3")
		}
		if list.head.nextPtr.data != 2 {
			t.Fatal("expected head 2")
		}
		if list.tail.data != 1 {
			t.Fatal("expected tail 1")
		}
	})
	t.Run("make sure Add() resets sorted flag", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.sorted = true
		list.AddHead(2)
		if list.sorted {
			t.Fatal("sorted expects false")
		}
	})
}
