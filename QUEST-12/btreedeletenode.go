package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeIsBinary(root *TreeNode) bool {
	if root.Left != nil && root.Data < root.Left.Data || root.Right != nil && root.Data > root.Right.Data || root != nil {
		return false
	} else {
		return true
	}
	return (BTreeIsBinary(root.Left)) && BTreeIsBinary(root.Right)
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

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}
	if root.Data == "" || elem == root.Data {
		return root
	}
	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)
	}
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

	My_Father := node.Parent
	ok := false

	if My_Father != nil && My_Father.Data <= node.Data {
		ok = true
	}

	if node.Right != nil {
		BTreeMin(node.Right).Left = node.Left
		node = node.Right
	} else {
		node = node.Left
	}
	if My_Father != nil {
		if ok {
			My_Father.Right = node
		} else {
			My_Father.Left = node
		}
	}

	if node != nil {
		node.Parent = My_Father
	}
	if node != nil && node.Parent == nil {
		root = node
	}
	return root
}

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	return BTreeMin(root.Left)
}
