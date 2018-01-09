package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	Variable   node.Node
	Expression node.Node
}
