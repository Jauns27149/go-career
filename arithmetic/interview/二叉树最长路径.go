/*
求二叉树中最大深度
*/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type example struct {
	root []int
	res  int
}

func main() {
	exs := []example{
		{[]int{1, 2, 2, 3, 0, 0, 5, 4, 0, 0, 4}, 4},
	}
	for _, ex := range exs {
		t := &TreeNode{}
		t.convert(ex.root)
		l := calculateDepth(t)
		fmt.Printf("计算：%v->答案：%v\n", l, ex.res)
	}
}

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 1
	}
	n := 0
	n = findEnd(root, n)
	return n
}

func findEnd(t *TreeNode, n int) int {
	if t.Left != nil || t.Right != nil {
		n++
	}
	ln, rn := n, n
	if t.Left != nil {
		ln = findEnd(t.Left, n)
	}
	if t.Right != nil {
		rn = findEnd(t.Right, n)
	}
	if ln > rn {
		return ln
	} else {
		return rn
	}
}

// 构建二叉树
func (t *TreeNode) convert(root []int) {
	if len(root) == 0 || root[0] == 0 {
		return
	}
	ch := make(chan *TreeNode, 5)
	ch <- t
	for i, v := range root {
		node := <-ch
		if v > 0 && i < len(root)-1 {
			node.Val = v
			node.Left = &TreeNode{}
			node.Right = &TreeNode{}
			ch <- node.Left
			ch <- node.Right
		}
		if v > 0 && i == len(root)-1 {
			node.Val = v
		}
	}
}

// levelOrderTraversal 进行层序遍历并返回结果切片
func (t *TreeNode) levelOrderTraversal() []int {
	if t == nil {
		return []int{}
	}
	var result []int
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}
