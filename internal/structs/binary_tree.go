package structs

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

// here we are dealing with nil receiver
func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func BinaryTree(show bool) {
	if show {
		var it *IntTree
		it = it.Insert(7)
		it = it.Insert(77)
		it = it.Insert(777)
		it = it.Insert(7777)

		fmt.Println(it.Contains(7))     // true
		fmt.Println(it.Contains(77777)) // false
	}
}
