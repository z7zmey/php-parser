package position

import (
	"fmt"

	"github.com/z7zmey/php-parser/node"
)

// Position represents node position
type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

func (p Position) String() string {
	return fmt.Sprintf("Pos{Line: %d-%d Pos: %d-%d}", p.StartLine, p.EndLine, p.StartPos, p.EndPos)
}

// Positions a collection of positions attached to nodes
type Positions map[node.Node]*Position

// AddPosition attaches a position to a node
func (p Positions) AddPosition(node node.Node, position *Position) {
	p[node] = position
}
