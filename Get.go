package genericlist

// Get - Return the data at the current node
func (list *LinkedList[T]) Get() T {
	return list.curr.data
}
