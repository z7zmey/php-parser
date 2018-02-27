package name

import (
	"github.com/z7zmey/php-parser/node"
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
