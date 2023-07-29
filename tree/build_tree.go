package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorder_tree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preorder_tree(root.Left)
	preorder_tree(root.Right)
}

// 利用数组创建一个二叉树
func buildTree(arr []string) *TreeNode {
	root_node := []*TreeNode{}

	for i := 0; i < len(arr)/2; i++ {
		node := new(TreeNode)
		node.Val, _ = strconv.Atoi(arr[i])

		if 2*i+1 < len(arr) && arr[2*i+1] != "null" {
			val, _ := strconv.Atoi(arr[2*i+1])
			node.Left = create_node(val)
		}
		if 2*i+2 < len(arr) && arr[2*i+2] != "null" {
			val, _ := strconv.Atoi(arr[2*i+2])
			node.Right = create_node(val)
		}

		root_node = append(root_node, node)
	}
	fmt.Println(len(root_node))
	return root_node[0]
}

// 创建子节点
func create_node(val int) *TreeNode {
	root := new(TreeNode)
	root.Val = val
	return root
}

func main() {
	arr := []string{"1", "2", "2", "3", "3", "null", "null", "4", "4"}
	root := buildTree(arr)
	preorder_tree(root)
}
