package tree

import (
	"reflect"
	"testing"
)

func TestAddChild(t *testing.T) {
	/*
	**
	**              1
	**    --------------------
	**    2         3        4
	**    |                  |
	**    5                  6
	**    |
	**    7
	 */
	root := Node{
		Id:       1,
		Value:    1,
		Children: []*Node{},
	}
	AddChild(&root, &Node{Id: 2, Value: 2})
	AddChild(&root, &Node{Id: 3, Value: 3})
	AddChild(&root, &Node{Id: 4, Value: 4})
	AddChild(root.Children[0], &Node{Id: 5, Value: 5})
	AddChild(root.Children[2], &Node{Id: 6, Value: 6})
	AddChild(root.Children[0].Children[0], &Node{Id: 7, Value: 7})
	expectedValues := []int{
		5,
		6,
		7,
	}
	actualValues := []int{
		root.Children[0].Children[0].Id,
		root.Children[2].Children[0].Id,
		root.Children[0].Children[0].Children[0].Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in AddChild. Expected: %d, got: %d", expectedValues, actualValues)
	}
}

func TestAddChildById(t *testing.T) {
	/*
	**
	**              1
	**    --------------------
	**    2         3        4
	**  -----       |
	**  5   6       7
	**  |       ---------
	**  8       9  10  11
	 */
	root := Node{
		Id:       1,
		Value:    1,
		Children: []*Node{},
	}
	AddChild(&root, &Node{Id: 2, Value: 2})
	AddChild(&root, &Node{Id: 3, Value: 3})
	AddChild(&root, &Node{Id: 4, Value: 4})
	AddChildById(&root, &Node{Id: 5, Value: 5}, 2)
	AddChildById(&root, &Node{Id: 6, Value: 6}, 2)
	AddChildById(&root, &Node{Id: 7, Value: 7}, 3)
	AddChildById(&root, &Node{Id: 8, Value: 8}, 5)
	AddChildById(&root, &Node{Id: 9, Value: 9}, 7)
	AddChildById(&root, &Node{Id: 10, Value: 10}, 7)
	AddChildById(&root, &Node{Id: 11, Value: 11}, 7)
	expectedValues := []int{
		8,
		6,
		11,
	}
	actualValues := []int{
		root.Children[0].Children[0].Children[0].Id,
		root.Children[0].Children[1].Id,
		root.Children[1].Children[0].Children[2].Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in AddChild. Expected: %d, got: %d", expectedValues, actualValues)
	}
}
