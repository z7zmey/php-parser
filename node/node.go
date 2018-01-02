package node

type Node interface {
	Attributer
	Positioner
	Name() string
	Walk(v Visitor)
}

type Attributer interface {
	Attributes() map[string]interface{}
	Attribute(key string) interface{}
	SetAttribute(key string, value interface{}) Node
}

type Positioner interface {
	Position() *Position
	SetPosition(p *Position) Node
}

type Position struct {
	StartLine int
	EndLine   int
}
