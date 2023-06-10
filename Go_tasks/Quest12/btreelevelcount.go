package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	countleft := BTreeLevelCount(root.Left)
	countright := BTreeLevelCount(root.Right)

	if countleft > countright {
		return countleft + 1
	} else {
		return countright + 1
	}
}
