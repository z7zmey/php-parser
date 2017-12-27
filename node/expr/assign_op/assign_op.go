package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	name       string
	variable   node.Node
	expression node.Node
}
