package orgmode

import (
	"encoding/json"
	"fmt"
)

type Tree struct {
	Nodes    []*Node `json:"nodes"`
	Subtrees []*Tree `json:"subtrees"`
	parent   *Tree
}

func NewTree(nodes []*Node) *Tree {
	tree := Tree{Nodes: nodes}
	tree.unflatten()

	return &tree
}

func (tree *Tree) addNode(node *Node) {
	node.parent = node.findParent(tree.Nodes)

	if node.Position == 0 {
		node.Position = 1
	}

	tree.Nodes = append(tree.Nodes, node)
}

func (tree *Tree) addSubtree(subtree *Tree) {
	subtree.parent = tree
	tree.Subtrees = append(tree.Subtrees, subtree)
}

func (tree Tree) isEmpty() bool {
	return len(tree.Nodes) == 0
}

func (tree Tree) lastNode() *Node {
	return tree.Nodes[len(tree.Nodes)-1]
}

func (tree *Tree) indexOfNode(searchNode *Node) int {
	for i, node := range tree.Nodes {
		if node == searchNode {
			return i
		}
	}

	return -1
}

func (tree *Tree) deleteNode(node *Node) {
	i := tree.indexOfNode(node)

	if i == -1 {
		return
	}

	if i == 0 {
		tree.Nodes = tree.Nodes[1:]
	} else if i == len(tree.Nodes)-1 {
		tree.Nodes = tree.Nodes[:len(tree.Nodes)-1]
	} else {
		tree.Nodes = append(tree.Nodes[:i], tree.Nodes[i+1:]...)
	}
}

func (tree *Tree) unflatten() {
	subtrees := getSubtrees(tree.Nodes)

	for _, s := range subtrees {
		tree.addSubtree(s)

		for _, n := range s.Nodes {
			tree.deleteNode(n)
		}
	}

	for _, subtree := range tree.Subtrees {
		subtree.unflatten()
	}
}

func getSubtrees(ns []*Node) []*Tree {

	if len(ns) == 1 {
		return []*Tree{}
	}

	root := ns[0]
	nodes := ns[1:]

	subtree := &Tree{Nodes: []*Node{root}}
	var subtrees []*Tree

	for _, node := range nodes {
		if node.Position > root.Position {
			subtree.addNode(node)
		} else {
			subtrees = append(subtrees, subtree)

			root = node

			subtree = &Tree{Nodes: []*Node{root}}

		}

		if node == nodes[len(nodes)-1] {
			subtrees = append(subtrees, subtree)
		}
	}

	if len(subtrees) > 1 {
		return subtrees
	} else {
		return getSubtrees(nodes)
	}
}

func (tree Tree) toJson() ([]byte, error) {
	json, err := json.Marshal(tree)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func printTree(tree Tree) {
	for _, node := range tree.Nodes {
		line := ""
		for i := 0; i < node.Position; i++ {
			line = line + "*"
		}

		line = line + " " + node.Headline
		fmt.Println(line)
	}

	for _, subtree := range tree.Subtrees {
		printTree(*subtree)
	}
}
