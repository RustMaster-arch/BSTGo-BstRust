package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

func newNode[T constraints.Ordered](value T) Node[T] {
	return Node[T]{
		left:  nil,
		right: nil,
		value: value,
	}
}

type Bst[T constraints.Ordered] struct {
	root   *Node[T]
	length uint16
}

func newBst[T constraints.Ordered](value T) Bst[T] {
	newNode := newNode(value)
	return Bst[T]{
		root: &newNode,
	}
}

func (bst *Bst[T]) insert(value T) {
	n := newNode(value)
	if value > bst.root.value && bst.root.right == nil {
		bst.root.right = &n
		bst.length++
		return
	}

	if value < bst.root.value && bst.root.left == nil {
		bst.root.left = &n
		bst.length++
		return
	}

	bst.InsertHelp(value)

}

func (bst *Bst[T]) InsertHelp(value T) {
	n := newNode(value)
	curr := bst.root
	for curr.left != nil || curr.right != nil {
		if value == curr.value {
			return
		}
		if value > curr.value {
			if curr.right != nil {
				curr = curr.right
				continue
			}
			break
		}
		if value < curr.value {
			if curr.left != nil {
				curr = curr.left
				continue
			}
			break
		}
	}
	if value > curr.value {
		curr.right = &n
	}

	if value < curr.value {
		curr.left = &n
	}
}

func (bst *Bst[T]) lookup(value T) T {
	curr := bst.root
	for curr.left != nil || curr.right != nil {
		if value == curr.value {
			break
		}
		if value > curr.value {
			if curr.right != nil {
				curr = curr.right
				continue
			}
			break
		}
		if value < curr.value {
			if curr.left != nil {
				curr = curr.left
				continue
			}
			break
		}
	}
	if value > curr.value {
		if curr.right != nil {
			return curr.right.value
		}
		fmt.Println("could not find value, returning nearest value to target instead")
	}

	if value < curr.value {
		if curr.left != nil {
			return curr.left.value
		}
		fmt.Println("could not find value, returning nearest value to target instead")
	}
	return curr.value
}

func main() {
	myBst := newBst(9)
	myBst.insert(20)
	myBst.insert(4)
	myBst.insert(6)
	myBst.insert(170)
	myBst.insert(15)
	myBst.insert(1)

	fmt.Println()

	fmt.Print("final output\n")
	myBst.Print()
	value := myBst.lookup(15)
	fmt.Println(value)
}
