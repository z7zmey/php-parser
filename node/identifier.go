package node

type Identifier struct {
	position *Position
	Value    string
}

func NewIdentifier(Value string) Node {
	return &Identifier{
		nil,
		Value,
	}
}

func (n Identifier) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n Identifier) Position() *Position {
	return n.position
}

func (n Identifier) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
