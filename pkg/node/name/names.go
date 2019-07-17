package name

import (
	"github.com/z7zmey/php-parser/pkg/node"
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
