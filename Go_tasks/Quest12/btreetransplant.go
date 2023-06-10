package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return rplc
	}

	if node == nil {
		return root
	}

	if node == root {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
