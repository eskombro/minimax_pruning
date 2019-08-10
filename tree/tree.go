package tree

import ()

type Node struct {
	Id            int
	Value         int
	Children      []*Node
	SelectedChild *Node
}

/*
** Add a new Node to the children list of a given parent address
 */
func AddChild(parent *Node, child *Node) {
	parent.Children = append(parent.Children, child)
}

/*
** Add a new Node to the children list of a given parent id.
** root is the starting point of the tree to look for parent
 */
func AddChildById(root *Node, child *Node, parentId int) int {
	if root.Id == parentId {
		AddChild(root, child)
		return 1
	} else {
		current := root
		for _, n := range current.Children {
			if n.Id == parentId {
				AddChild(n, child)
				return 1
			}
		}
		for _, n := range current.Children {
			ret := AddChildById(n, child, parentId)
			if ret == 1 {
				return 1
			}
		}
	}
	return 0
}
