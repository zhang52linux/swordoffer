package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 遍历方式，dfs左中右
func kthLargest(root *TreeNode, k int) (ans int) {
	i := 0
	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		dfs(n.Right)
		i++
		if i == k {
			ans = n.Val
			return
		}
		dfs(n.Left)
	}
	dfs(root)
	return
}

func main() {
	root := new(TreeNode)
	left := new(TreeNode)
	right := new(TreeNode)
	root.Val = 1
	left.Val = 0
	right.Val = 3
	root.Left = left
	root.Right = right
	res := kthLargest(root, 1)
	fmt.Println(res)
}
