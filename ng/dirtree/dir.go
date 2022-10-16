package dirtree

import (
	"os"
)

type DirNode struct {
	name     string
	path     string
	parent   *DirNode
	isDir    bool
	cached   bool
	children []*DirNode
}

func NewDirNode(name string, parent *DirNode) *DirNode {
	node := new(DirNode)

	node.name = name
	node.parent = parent
	node.path = ""
	node.cached = false

	if parent != nil {
		node.path += parent.path

		if parent.parent != nil {
			node.path += "/"
		}
	}
	node.path += name

	f, err := os.Open(node.path)
	if err != nil {
		// in case it does not exist
		return nil
	}

	defer f.Close()

	s, err := f.Stat()
	node.isDir = err == nil && s.IsDir()

	return node
}

func (node *DirNode) FetchChildren() {
	if node.isDir {
		node.fetchChildren()
	}
}

func (node *DirNode) fetchChildren() {
	descriptor, err := os.Open(node.path)
	if err != nil {
		return
	}

	content, err := descriptor.Readdirnames(0)
	if err != nil {
		return
	}

	for _, name := range content {
		n := NewDirNode(name, node)
		if n != nil {
			node.children = append(node.children, n)
		}
	}

	node.cached = true
}

func (node *DirNode) IsDir() bool {
	return node.isDir
}

func (node *DirNode) GetParent() *DirNode {
	if node == nil {
		return nil
	}

	return node.parent
}

func (node *DirNode) GetChildren() []*DirNode {
	if node == nil {
		return []*DirNode{}
	}

	if !node.cached {
		node.FetchChildren()
	}

	if len(node.children) <= 0 {
		return []*DirNode{}
	}

	return node.children
}

func (node *DirNode) GetChildrenNames() []string {
	names := []string{}

	if node == nil {
		return names
	}

	if !node.cached {
		node.FetchChildren()
	}

	if len(node.children) <= 0 {
		return names
	}

	for _, n := range node.children {
		names = append(names, n.name)
	}

	return names
}

func (node *DirNode) GetPath() string {
	return node.path
}

func (node *DirNode) GetName() string {
	return node.name
}

func (node *DirNode) GetChildByName(name string) *DirNode {
	if !node.cached {
		node.FetchChildren()
	}

	if node == nil || len(node.children) <= 0 {
		return nil
	}

	for _, child := range node.children {
		if child.name == name {

			if child.IsDir() {
				return child
			} else {
				break
			}
		}
	}

	return nil
}
