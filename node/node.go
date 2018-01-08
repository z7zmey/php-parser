package node

import "fmt"

type Node interface {
	Positioner
	Attributes() map[string]interface{}
	Walk(v Visitor)
}

type Positioner interface {
	Position() *Position
	SetPosition(p *Position) Node
}

type Position struct {
	StartLine int
	EndLine   int
	StartPos  int
	EndPos    int
}

func (p Position) String() string {
	return fmt.Sprintf("Pos{Line: %d-%d Pos: %d-%d}", p.StartLine, p.EndLine, p.StartPos, p.EndPos)
}
