package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	position   *node.Position
	Variable   node.Node
	Expression node.Node
}
