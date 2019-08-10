package minimax

import (
	"reflect"
	"testing"

	tr "github.com/gogogomoku/minimax_pruning/tree"
)

func TestLaunchMinimax(t *testing.T) {
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
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1

	root := tr.Node{
		Id:       1,
		Value:    MinInt,
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
	LaunchMinimax(&root, 3)
	expectedValues := []int{
		2,
		4,
		9,
	}
	actualValues := []int{
		root.SelectedChild.Id,
		root.SelectedChild.SelectedChild.Id,
		root.SelectedChild.SelectedChild.SelectedChild.Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in AddChild. Expected: %d, got: %d", expectedValues, actualValues)
	}

	/*
	 **
	 **                                1(8)
	 **                ---------------------------------
	 **                2(8)                               3(4)
	 **        -----------------              -----------------------
	 **        4(10)        5(8)              6(6)                7(4)
	 ** -------------    ------------   ----------------      ---------------
	 ** 8(10)    9(9)    10(8)  11(7)   12(6)    13(5)PR      14(4)     15(3)
	 */
	root = tr.Node{
		Id:       1,
		Value:    MinInt,
		Children: []*tr.Node{},
	}
	tr.AddChild(&root, &tr.Node{Id: 2, Value: 0})
	tr.AddChildById(&root, &tr.Node{Id: 4, Value: 0}, 2)
	tr.AddChildById(&root, &tr.Node{Id: 5, Value: 0}, 2)
	tr.AddChildById(&root, &tr.Node{Id: 8, Value: 10}, 4)
	tr.AddChildById(&root, &tr.Node{Id: 9, Value: 9}, 4)
	tr.AddChildById(&root, &tr.Node{Id: 10, Value: 8}, 5)
	tr.AddChildById(&root, &tr.Node{Id: 11, Value: 7}, 5)
	tr.AddChild(&root, &tr.Node{Id: 3, Value: 0})
	tr.AddChildById(&root, &tr.Node{Id: 6, Value: 0}, 3)
	tr.AddChildById(&root, &tr.Node{Id: 7, Value: 0}, 3)
	tr.AddChildById(&root, &tr.Node{Id: 12, Value: 6}, 6)
	tr.AddChildById(&root, &tr.Node{Id: 13, Value: 5}, 6)
	tr.AddChildById(&root, &tr.Node{Id: 14, Value: 4}, 7)
	tr.AddChildById(&root, &tr.Node{Id: 15, Value: 3}, 7)
	LaunchMinimax(&root, 3)
	expectedValues = []int{
		2,
		5,
		10,
	}
	actualValues = []int{
		root.SelectedChild.Id,
		root.SelectedChild.SelectedChild.Id,
		root.SelectedChild.SelectedChild.SelectedChild.Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in Minimax. Expected: %d, got: %d", expectedValues, actualValues)
	}
}
