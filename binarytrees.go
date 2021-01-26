/*
*  Gotour exercise: Equivalent Binary Trees
 */

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		ch <- t.Value
		Walk(t.Left, ch)
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for xElem := range ch1 {
		xMap[xElem]++
	}
	for yElem := range ch2 {
		yMap[yElem]++
	}

	for xMapKey, xMapVal := range xMap {
		if yMap[xMapKey] != xMapVal {
			return false
		}
	}
	return true
}

func testWalk() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}

func testSame() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

func main() {
	defer fmt.Println("bye")
	testWalk()
	testSame()
}
