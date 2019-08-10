package minimax

import (
	"fmt"

	tr "github.com/gogogomoku/minimax_pruning/tree"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

var nodeCounter int

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxChild(node *tr.Node, depth int, max bool) int {
	bestValue := MinInt
	for _, ch := range node.Children {
		returnValue := MinimaxRecursive(ch, depth-1, !max)
		bestValue = maximum(bestValue, returnValue)
		if node.SelectedChild == nil || ch.Value > node.SelectedChild.Value {
			node.SelectedChild = ch
		}
	}
	return bestValue
}

func getMinChild(node *tr.Node, depth int, max bool) int {
	bestValue := MaxInt
	for _, ch := range node.Children {
		returnValue := MinimaxRecursive(ch, depth-1, !max)
		bestValue = minimum(bestValue, returnValue)
		if node.SelectedChild == nil || ch.Value < node.SelectedChild.Value {
			node.SelectedChild = ch
		}
	}
	return bestValue
}

func MinimaxRecursive(node *tr.Node, depth int, max bool) int {
	nodeCounter++
	if depth == 0 || len(node.Children) == 0 {
		return node.Value
	}
	if max {
		node.Value = getMaxChild(node, depth, max)
		return node.Value
	} else {
		node.Value = getMinChild(node, depth, max)
		return node.Value
	}
}

func LaunchMinimax(graph *tr.Node, depth int) {
	nodeCounter = 0
	fmt.Println("Launching Minimax with depth ", depth)
	graph.Value = MinimaxRecursive(graph, depth, true)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.Value)
	fmt.Println("Nodes checked:", nodeCounter)
}
