package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func WalkAux(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	WalkAux(t.Left, ch)
	ch <-t.Value
	WalkAux(t.Right, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	WalkAux(t, ch)
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
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			return false
		case ok1 == false:
			return true
		case v1 != v2:
			return false
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
