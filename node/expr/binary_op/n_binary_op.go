package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

// BinaryOp node
type BinaryOp struct {
	Left  node.Node
	Right node.Node
}
