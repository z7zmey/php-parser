package position

import (
	"fmt"
)

// Position represents node position
type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

// NewPosition Position constructor
func NewPosition(StartLine int, EndLine int, StartPos int, EndPos int) *Position {
	return &Position{
		StartLine: StartLine,
		EndLine:   EndLine,
		StartPos:  StartPos,
		EndPos:    EndPos,
	}
}

func (p Position) String() string {
	return fmt.Sprintf("Pos{Line: %d-%d Pos: %d-%d}", p.StartLine, p.EndLine, p.StartPos, p.EndPos)
}
