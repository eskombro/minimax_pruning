package main

import (
	"fmt"
	gr "github.com/eskombro/minimax/graph"
	mm "github.com/eskombro/minimax/minimax"
)

func createGraph() *gr.Node {
	/*
	**
	**                                1
	**                ---------------------------------
	**                2                               3
	**        -----------------              ----------------
	**        4               5              6              7
	** -------------    ------------   --------------  ------------
	** 8(-1)    9(3)    10(5)  11(1)   12(-6)  13(-4)  14(0)  15(9)
	 */
	root := gr.Node{
		Id:       1,
		Value:    0,
		Children: []*gr.Node{},
	}
	gr.AddChild(&root, &gr.Node{Id: 2, Value: 0})
	gr.AddChildById(&root, &gr.Node{Id: 4, Value: 0}, 2)
	gr.AddChildById(&root, &gr.Node{Id: 5, Value: 0}, 2)
	gr.AddChildById(&root, &gr.Node{Id: 8, Value: -1}, 4)
	gr.AddChildById(&root, &gr.Node{Id: 9, Value: 3}, 4)
	gr.AddChildById(&root, &gr.Node{Id: 10, Value: 5}, 5)
	gr.AddChildById(&root, &gr.Node{Id: 11, Value: 1}, 5)
	gr.AddChild(&root, &gr.Node{Id: 3, Value: 0})
	gr.AddChildById(&root, &gr.Node{Id: 6, Value: 0}, 3)
	gr.AddChildById(&root, &gr.Node{Id: 7, Value: 0}, 3)
	gr.AddChildById(&root, &gr.Node{Id: 12, Value: -6}, 6)
	gr.AddChildById(&root, &gr.Node{Id: 13, Value: -4}, 6)
	gr.AddChildById(&root, &gr.Node{Id: 14, Value: 0}, 7)
	gr.AddChildById(&root, &gr.Node{Id: 15, Value: 9}, 7)
	return &root
}

func main() {
	graph := createGraph()
	mm.LaunchMinimax(graph, 2)
	current := graph
	for current.SelectedChild != nil {
		fmt.Println(current.Id)
		current = current.SelectedChild
	}
	fmt.Println(current.Id)
}
