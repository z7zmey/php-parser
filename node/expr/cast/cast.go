package cast

import (
	"github.com/z7zmey/php-parser/node"
)

type Cast struct {
	attributes map[string]interface{}
	position   *node.Position
	expr       node.Node
}
