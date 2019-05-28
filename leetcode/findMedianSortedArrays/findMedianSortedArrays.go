package findMedianSortedArrays

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var rootNode *TNode

	// fmt.Println("----------", nums1, nums2)

	for i, _ := range nums1 {
		rootNode = insert(rootNode, nums1[i])
	}
	for i, _ := range nums2 {
		rootNode = insert(rootNode, nums2[i])
	}
	if rootNode == nil {
		fmt.Println("rootNode == nil")
		return 0.0
	}
	fmt.Println("----------", rootNode.Key, rootNode.Left.Key, rootNode.Right.Key)
	fmt.Println("----------", rootNode.Right.Left.Key, rootNode.Right.Right.Key)
	fmt.Println("----------height", height(rootNode), height(rootNode.Left), height(rootNode.Right))

	if height(rootNode.Left) == height(rootNode.Right) {
		return float64(rootNode.Key)
	}
	if height(rootNode.Left) > height(rootNode.Right) {
		return float64(rootNode.Key+rootNode.Left.Key) / 2
	}
	if height(rootNode.Left) < height(rootNode.Right) {
		return float64(rootNode.Key+rootNode.Right.Key) / 2
	}

	return 0.0
}

type TNode struct {
	Key    int
	Height int
	Left   *TNode
	Right  *TNode
}

func insert(n *TNode, key int) *TNode {
	if n == nil {
		return &TNode{
			Key: key,
		}
	}
	if key < n.Key {
		n.Left = insert(n.Left, key)
	} else if key >= n.Key {
		n.Right = insert(n.Right, key)
	}

	n.Height = 1 + int(math.Max(float64(height(n.Left)), float64(height(n.Right))))

	balance := getBalance(n)
	switch {
	case balance > 1 && key < n.Left.Key:
		n.Height -= 2
		return rightRotate(n)
	case balance < -1 && key >= n.Right.Key:
		n.Height -= 2
		fmt.Println("leftRotate", n.Key)
		return leftRotate(n)
	case balance > 1 && key >= n.Left.Key:
		n.Left.Height -= 1
		n.Height -= 2
		n.Left.Right.Height += 1
		n.Left = leftRotate(n.Left)
		return rightRotate(n)
	case balance < -1 && key < n.Right.Key:
		n.Right.Height -= 1
		n.Height -= 2
		n.Right.Left.Height += 1
		n.Right = rightRotate(n.Right)
		return leftRotate(n)
	}
	return n
}

func rightRotate(n *TNode) *TNode {
	rnode := n.Left
	rightNode := n.Left.Right
	n.Left.Right = n
	n.Left = rightNode
	return rnode
}

func leftRotate(n *TNode) *TNode {
	rnode := n.Right
	leftNode := n.Right.Left
	n.Right.Left = n
	n.Right = leftNode
	return rnode
}

func getBalance(n *TNode) int {
	var leftHeight, rightHeight int
	leftHeight = height(n.Left) + 1
	rightHeight = height(n.Right) + 1
	return leftHeight - rightHeight
}

func height(n *TNode) int {
	if n == nil {
		return -1
	}
	return n.Height
}
