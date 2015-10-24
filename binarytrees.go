package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
//func Walk(t *tree.Tree, ch chan int) {
//	if t == nil {
//		return
//	}
//	if t.Left == nil && t.Right == nil {
//		ch <- t.Value
//		return
//	}
//	Walk(t.Left, ch)
//	ch <- t.Value
//	Walk(t.Right, ch)
//	return
//}

func walkImpl(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkImpl(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkImpl(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkImpl(t, ch)
	// Need to close the channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		x, ok1 := <-ch1
		y, ok2 := <-ch2
		if x != y || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

