package node

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
}
