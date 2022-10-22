package tree

import (
	"github.com/TheAlgorithms/Go/constraints"
	"github.com/TheAlgorithms/Go/math/max"
)

type AVL[T constraints.Ordered] struct {
	*binaryTree[T]
}

// NewAVL create a novel AVL tree
func NewAVL[T constraints.Ordered]() *AVL[T] {
	var leaf *Node[T]
	return &AVL[T]{
		binaryTree: &binaryTree[T]{
			Root: leaf,
			NIL:  leaf,
		},
	}
}

// Push a chain of Node's into the AVL Tree
func (avl *AVL[T]) Push(keys ...T) {
	for _, k := range keys {
		avl.Root = avl.pushHelper(avl.Root, k)
	}
}

// Delete a Node from the AVL Tree
func (avl *AVL[T]) Delete(key T) bool {
	if !avl.Has(key) {
		return false
	}

	avl.Root = avl.deleteHelper(avl.Root, key)
	return true
}

func (avl *AVL[T]) pushHelper(root *Node[T], key T) *Node[T] {
	if avl.isNil(root) {
		return &Node[T]{
			Key:    key,
			Left:   avl.NIL,
			Right:  avl.NIL,
			Parent: avl.NIL,
			Height: 1,
		}
	}

	switch {
	case key < root.Key:
		root.Left = avl.pushHelper(root.Left, key)
	case key > root.Key:
		root.Right = avl.pushHelper(root.Right, key)
	default:
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	if bFactor > 1 {
		switch {
		case key < root.Left.Key:
			return avl.rightRotate(root)
		case key > root.Left.Key:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	}

	if bFactor < -1 {
		switch {
		case key > root.Right.Key:
			return avl.leftRotate(root)
		case key < root.Right.Key:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVL[T]) deleteHelper(root *Node[T], key T) *Node[T] {
	if avl.isNil(root) {
		return root
	}

	switch {
	case key < root.Key:
		root.Left = avl.deleteHelper(root.Left, key)
	case key > root.Key:
		root.Right = avl.deleteHelper(root.Right, key)
	default:
		if avl.isNil(root.Left) || avl.isNil(root.Right) {
			tmp := root.Left
			if !avl.isNil(root.Left) {
				tmp = root.Right
			}

			if avl.isNil(tmp) {
				root = avl.NIL
			} else {
				*root = *tmp
			}
		} else {
			tmp := avl.minimum(root.Right)
			root.Key = tmp.Key
			root.Right = avl.deleteHelper(root.Right, tmp.Key)
		}
	}

	if avl.isNil(root) {
		return root
	}

	// balance the tree
	root.Height = avl.height(root)
	bFactor := avl.balanceFactor(root)
	switch {
	case bFactor > 1:
		switch {
		case avl.balanceFactor(root.Left) >= 0:
			return avl.rightRotate(root)
		default:
			root.Left = avl.leftRotate(root.Left)
			return avl.rightRotate(root)
		}
	case bFactor < -1:
		switch {
		case avl.balanceFactor(root.Right) <= 0:
			return avl.leftRotate(root)
		default:
			root.Right = avl.rightRotate(root.Right)
			return avl.leftRotate(root)
		}
	}

	return root
}

func (avl *AVL[T]) height(root *Node[T]) int {
	if avl.isNil(root) {
		return 0
	}

	var leftHeight, rightHeight int
	if !avl.isNil(root.Left) {
		leftHeight = root.Left.Height
	}
	if !avl.isNil(root.Right) {
		rightHeight = root.Right.Height
	}
	return 1 + max.Int(leftHeight, rightHeight)
}

// balanceFactor : negative balance factor means subtree Root is heavy toward Left
// and positive balance factor means subtree Root is heavy toward Right side
func (avl *AVL[T]) balanceFactor(root *Node[T]) int {
	var leftHeight, rightHeight int
	if !avl.isNil(root.Left) {
		leftHeight = root.Left.Height
	}
	if !avl.isNil(root.Right) {
		rightHeight = root.Right.Height
	}
	return leftHeight - rightHeight
}

func (avl *AVL[T]) leftRotate(root *Node[T]) *Node[T] {
	y := root.Right
	yl := y.Left
	y.Left = root
	root.Right = yl

	root.Height = avl.height(root)
	y.Height = avl.height(y)
	return y
}

func (avl *AVL[T]) rightRotate(root *Node[T]) *Node[T] {
	y := root.Left
	yr := y.Right
	y.Right = root
	root.Left = yr

	root.Height = avl.height(root)
	y.Height = avl.height(y)
	return y
}
