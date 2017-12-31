package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	name       string
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}
