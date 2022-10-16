package dirtree

type DirTree struct {
	Root *DirNode
}

func New() *DirTree {
	tree := new(DirTree)

	tree.Root = NewDirNode("/", nil)
	tree.Root.FetchChildren()

	return tree
}
