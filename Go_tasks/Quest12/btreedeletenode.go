package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		// Case 1: No child or one child
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}

		// Case 2: Two children
		// Find the minimum value in the right subtree (successor)
		successor := BTreeMinValue(root.Right)
		// Copy the successor's data to the current node
		root.Data = successor.Data
		// Delete the successor from the right subtree
		root.Right = BTreeDeleteNode(root.Right, successor)
	}

	return root
}

func BTreeMinValue(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}
