package ng

import (
	"os/user"

	"github.com/leo-alvarenga/gfile/ng/dirtree"
)

type LocationCursor struct {
	current *dirtree.DirNode
}

func New() *LocationCursor {
	cursor := new(LocationCursor)
	cursor.current = FILE_TREE.Root

	cursor.EnterChild("home")
	if cursor.current != FILE_TREE.Root {
		you, err := user.Current()

		if err == nil {
			cursor.EnterChild(you.Username)
		}
	}

	return cursor
}

func (cursor *LocationCursor) IsThisRoot() bool {
	return cursor.current.GetParent() == nil
}

func (cursor *LocationCursor) DangerouslySetCurrent(current *dirtree.DirNode) {
	cursor.current = current
}

func (cursor *LocationCursor) GetInfo() *dirtree.DirNode {
	return cursor.current
}

func (cursor *LocationCursor) EnterChild(name string) {
	child := cursor.current.GetChildByName(name)

	if child != nil {
		cursor.current = child
	}
}

func (cursor *LocationCursor) GoBack() {
	if cursor.current.GetParent() == nil {
		return
	}

	cursor.current = cursor.current.GetParent()
}
