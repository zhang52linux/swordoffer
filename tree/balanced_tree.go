package main

import "fmt"

//注：如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据要求：我们可以知道我们需要一个dfs函数，求出左子树的深度和右子树的深度
// 再把它们相减，如果值大于1则不是平衡二叉树
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return false
	}
	left_depth := max_depth(root.Left)
	right_depth := max_depth(root.Right)

	if right_depth-left_depth <= 1 {
		return true
	}
	return false
}

func max_depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(max_depth(root.Left), max_depth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := new(TreeNode)
	left := new(TreeNode)
	right := new(TreeNode)
	root.Val = 1
	root.Left = left
	root.Right = right
	res := isBalanced(root)
	fmt.Println(res)
}
