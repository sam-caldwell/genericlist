package genericlist

import (
	"testing"
)

func TestLinkedListNode_struct(t *testing.T) {
	_ = LinkedListNode[int]{
		data:    0,
		nextPtr: &LinkedListNode[int]{},
		prevPtr: &LinkedListNode[int]{},
	}
}
