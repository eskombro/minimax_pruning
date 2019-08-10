package main

import (
	"fmt"
	mm "github.com/gogogomoku/minimax_pruning/minimax"
	mmp "github.com/gogogomoku/minimax_pruning/pruning"
	tr "github.com/gogogomoku/minimax_pruning/tree"
)

func createTree() *tr.Node {
	/*
	**
	**                                1(3)**
	**                ---------------------------------
	**                2(3)**                      3(-4)
	**        -----------------              ----------------
	**        4(3)**       5(5)              6(-4)       7(9)
	** -------------    ------------   --------------  ------------
	** 8(-1)    9(3)**  10(5)  11(1)   12(-6)  13(-4)  14(0)  15(9)
	 */
	root := tr.Node{
		Id:       1,
		Value:    0,
		Children: []*tr.Node{},
	}
	tr.AddChild(&root, &tr.Node{Id: 2, Value: 0})
	tr.AddChildById(&root, &tr.Node{Id: 4, Value: 0}, 2)
	tr.AddChildById(&root, &tr.Node{Id: 5, Value: 0}, 2)
	tr.AddChildById(&root, &tr.Node{Id: 8, Value: -1}, 4)
	tr.AddChildById(&root, &tr.Node{Id: 9, Value: 3}, 4)
	tr.AddChildById(&root, &tr.Node{Id: 10, Value: 5}, 5)
	tr.AddChildById(&root, &tr.Node{Id: 11, Value: 1}, 5)
	tr.AddChild(&root, &tr.Node{Id: 3, Value: 0})
	tr.AddChildById(&root, &tr.Node{Id: 6, Value: 0}, 3)
	tr.AddChildById(&root, &tr.Node{Id: 7, Value: 0}, 3)
	tr.AddChildById(&root, &tr.Node{Id: 12, Value: -6}, 6)
	tr.AddChildById(&root, &tr.Node{Id: 13, Value: -4}, 6)
	tr.AddChildById(&root, &tr.Node{Id: 14, Value: 0}, 7)
	tr.AddChildById(&root, &tr.Node{Id: 15, Value: 9}, 7)
	return &root
}

func main() {
	tree := createTree()
	mm.LaunchMinimax(tree, 3)
	current := tree
	for current.SelectedChild != nil {
		fmt.Println(current.Id)
		current = current.SelectedChild
	}
	fmt.Println(current.Id)
	mmp.LaunchMinimaxPruning(tree, 3)
	current = tree
	for current.SelectedChild != nil {
		fmt.Println(current.Id)
		current = current.SelectedChild
	}
	fmt.Println(current.Id)
}
