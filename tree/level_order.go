package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 双队列实现从上到下打印二叉树。leetcode一遍过(2023.05.24) */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	front := make([]*TreeNode, 0)
	front = append(front, root)

	res := [][]int{}
	floor := 0

	for len(front) > 0 {
		rear := make([]*TreeNode, 0)
		tmp := []int{}

		for i := 0; i < len(front); i++ {
			tmp = append(tmp, front[i].Val)

			if front[i].Left != nil {
				rear = append(rear, front[i].Left)
			}
			if front[i].Right != nil {
				rear = append(rear, front[i].Right)
			}
		}
		res = append(res, tmp)
		front = rear
		floor++
	}
	return res
}

/* 与上面一题的区别在于：返回的数据结构不同，前面的为二维数组，后面的是一维数组 */
func levelOrder1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	front := make([]*TreeNode, 0)
	front = append(front, root)

	res := []int{}
	floor := 0

	for len(front) > 0 {
		rear := make([]*TreeNode, 0)

		for i := 0; i < len(front); i++ {
			res = append(res, front[i].Val)

			if front[i].Left != nil {
				rear = append(rear, front[i].Left)
			}
			if front[i].Right != nil {
				rear = append(rear, front[i].Right)
			}
		}
		front = rear
		floor++
	}
	return res
}

func main() {
	root := new(TreeNode)
	left := new(TreeNode)
	right := new(TreeNode)
	root.Val = 1
	root.Left = left
	root.Right = right
	res := levelOrder1(root)
	fmt.Println(res)
}
