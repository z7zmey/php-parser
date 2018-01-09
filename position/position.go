package position

import (
	"fmt"

	"github.com/z7zmey/php-parser/node"
)

type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

func (p Position) String() string {
	return fmt.Sprintf("Pos{Line: %d-%d Pos: %d-%d}", p.StartLine, p.EndLine, p.StartPos, p.EndPos)
}

type Positions map[node.Node]*Position

func (p Positions) AddPosition(node node.Node, position *Position) {
	p[node] = position
}
