package genericlist

// Prev - Move to the previous node, return nil if none.
func (list *LinkedList[T]) Prev() bool {
	list.lock.Lock()
	defer list.lock.Unlock()

	return list.movePrev()
}

func (list *LinkedList[T]) movePrev() bool {
	if list.curr.prevPtr != nil {
		list.curr = list.curr.prevPtr
		return true
	}
	return false
}
