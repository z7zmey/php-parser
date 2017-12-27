package binary_op

import (
	"github.com/z7zmey/php-parser/node"
)

type BinaryOp struct {
	name  string
	left  node.Node
	right node.Node
}
