package orgmode

func TreeFromFile(orgPath string) *Tree {
	return NewTree(nodesFromFile(orgPath))
}
