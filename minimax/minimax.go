package minimax

import (
	"fmt"

	gr "github.com/eskombro/minimax/graph"
)

func maxValuesFromChildren(node *gr.Node) {
	selected := 0
	for i, ch := range node.Children {
		if ch.UpdatedValue > node.Children[selected].UpdatedValue {
			selected = i
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func minValuesFromChildren(node *gr.Node) {
	selected := 0
	for i, ch := range node.Children {
		if ch.UpdatedValue < node.Children[selected].UpdatedValue {
			selected = i
		}
	}
	node.UpdatedValue = node.Children[selected].UpdatedValue
	node.SelectedChild = node.Children[selected]
}

func MinimaxRecursive(graph *gr.Node, depth int, max bool, start bool) {
	for _, node := range graph.Children {
		if depth >= 0 {
			MinimaxRecursive(node, depth-1, !max, false)
			if len(node.Children) == 0 || depth == 0 {
				node.UpdatedValue = node.Value

			} else {
				if max {
					maxValuesFromChildren(node)
				} else {
					minValuesFromChildren(node)
				}
			}
			fmt.Println("==========================")
			fmt.Println("  Depth:", depth, "Node:", node.Id)
			fmt.Println("  UpdatedValue:", node.UpdatedValue)
		}
	}
	if start {
		maxValuesFromChildren(graph)
	}
}

func LaunchMinimax(graph *gr.Node, depth int) {
	fmt.Println("Launching Minimax with depth ", depth)
	graph.UpdatedValue = graph.Value
	MinimaxRecursive(graph, depth, false, true)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.UpdatedValue)
}
