package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先搜索，最大深度等于左子树的深度与右子树的深度求max + 1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 广度优先搜索
/*
1. 使用两个队列，一个队列存放当前的层的元素，另一个队列存放下一层的元素
2. 使用一个队列，利用一个元素sz，记录当前层的个数，每次取出队列头结点，将该结点的左右结点往队列后追加
*/
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0

	for len(queue) > 0 {

		// 获取当前层的个数
		sz := len(queue)

		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz -= 1
		}
		ans += 1
	}
	return ans
}

func main() {
	root := new(TreeNode)
	left := new(TreeNode)
	right := new(TreeNode)
	root.Val = 1
	root.Left = left
	root.Right = right
	res := maxDepth1(root)
	fmt.Println(res)
}
