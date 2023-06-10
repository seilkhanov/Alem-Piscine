package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newnode := &NodeI{
		Data: data_ref,
		Next: nil,
	}

	// If newnode is inserted as the new head
	if l == nil || data_ref < l.Data {
		newnode.Next = l
		return newnode
	}

	// If newnode is inserted in the middle or at the end
	current := l
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}
	newnode.Next = current.Next
	current.Next = newnode

	return l
}
