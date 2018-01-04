package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BinaryOp struct {
	attributes map[string]interface{}
	position   *node.Position
	left       node.Node
	right      node.Node
}
