package node

type Nullable struct {
	attributes map[string]interface{}
	position   *Position
	expr       Node
}

func NewNullable(Expression Node) Node {
	return &Nullable{
		map[string]interface{}{},
		nil,
		Expression,
	}
}

func (n Nullable) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Nullable) Attribute(Key string) interface{} {
	return n.attributes[Key]
}

func (n Nullable) SetAttribute(Key string, Value interface{}) Node {
	n.attributes[Key] = Value
	return n
}

func (n Nullable) Position() *Position {
	return n.position
}

func (n Nullable) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Nullable) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
