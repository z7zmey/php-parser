package parser

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
)

// Positions a collection of positions attached to nodes
type Positions map[node.Node]*position.Position

// AddPosition attaches a position to a node
func (p Positions) AddPosition(node node.Node, position *position.Position) {
	p[node] = position
}
