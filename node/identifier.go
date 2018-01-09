package node

type Identifier struct {
	Value string
}

func NewIdentifier(Value string) *Identifier {
	return &Identifier{
		Value,
	}
}

func (n *Identifier) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *Identifier) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
