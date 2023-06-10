package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !BTreeIsBinary(root.Left) || !BTreeIsBinary(root.Right) {
		return false
	}

	if root.Left != nil && root.Right != nil && root.Left.Data > root.Right.Data {
		return false
	} else {
		return true
	}
}
