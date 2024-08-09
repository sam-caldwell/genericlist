package genericlist

import (
	"testing"
)

func TestLinkedList_struct(t *testing.T) {
	list := LinkedList[int]{
		head:  &LinkedListNode[int]{},
		curr:  &LinkedListNode[int]{},
		tail:  &LinkedListNode[int]{},
		count: 0,
	}
	list.lock.Lock()
	list.lock.Unlock()
}
