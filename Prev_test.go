package genericlist

import "testing"

func TestLinkedList_Prev(t *testing.T) {
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
	list.curr = list.tail
	for i := 10; i >= 1; i-- {
		if list.curr.data != i {
			t.Fatalf("got %d, want %d", list.curr.data, i)
		}
		list.Prev()
	}
}
