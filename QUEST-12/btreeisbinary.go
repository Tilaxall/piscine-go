package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	return (BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right))
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
