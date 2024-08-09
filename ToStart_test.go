package genericlist

import (
	"testing"
)

func TestLinkedList_ToStart(t *testing.T) {
	n := LinkedList[int]{}
	n.AddTail(1)
	n.AddTail(2)
	n.AddTail(3)
	n.AddTail(4)
	n.curr = n.tail
	if n.curr.data != 4 {
		t.Fatal("expected 4")
	}
	n.ToStart()
	if n.curr.data != 1 {
		t.Fatal("expected 1")
	}
}
