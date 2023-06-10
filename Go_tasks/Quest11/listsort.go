package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	// Split the linked list into two halves
	var previous *NodeI = nil
	slow := l
	fast := l

	for fast != nil && fast.Next != nil {
		previous = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	previous.Next = nil

	// Recursively sort the two halves
	left := ListSort(l)
	right := ListSort(slow)

	// Merge the sorted halves
	return merge(left, right)
}

func merge(left, right *NodeI) *NodeI {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	var result *NodeI
	if left.Data <= right.Data {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}
	return result
}
