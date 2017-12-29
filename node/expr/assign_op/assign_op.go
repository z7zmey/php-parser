package assign_op

import (
	"github.com/z7zmey/php-parser/node"
)

type AssignOp struct {
	name       string
	attributes map[string]interface{}
	variable   node.Node
	expression node.Node
}
