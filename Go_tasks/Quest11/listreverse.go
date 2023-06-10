package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListReverse(l *List) {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	var previous *NodeL = nil
	current := l.Head
	var next *NodeL = nil

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	l.Head, l.Tail = l.Tail, l.Head
}
