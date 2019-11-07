package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	hight := BTreeLevelCount(root)
	for i := 0; i < hight; i++ {
		Move(root, f, i)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	hleft, hright := 0, 0
	if root == nil {
		return 0
	}

	hleft, hright = BTreeLevelCount(root.Left), BTreeLevelCount(root.Right)

	if hleft > hright {
		return hleft + 1
	} else {
		return hright + 1
	}
}

func Move(root *TreeNode, f func(...interface{}) (int, error), hight int) {
	if root == nil {
		return
	}
	if hight == 0 {
		f(root.Data)
	} else if hight > 0 {
		Move(root.Left, f, hight-1)
		Move(root.Right, f, hight-1)
	}
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
