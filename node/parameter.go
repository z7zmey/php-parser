package node

type Parameter struct {
	attributes   map[string]interface{}
	position     *Position
	VariableType Node
	Variable     Node
	DefaultValue Node
}

func NewParameter(VariableType Node, Variable Node, DefaultValue Node, byRef bool, variadic bool) Node {
	return &Parameter{
		map[string]interface{}{
			"byRef":    byRef,
			"variadic": variadic,
		},
		nil,
		VariableType,
		Variable,
		DefaultValue,
	}
}

func (n Parameter) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Parameter) Attribute(Key string) interface{} {
	return n.attributes[Key]
}

func (n Parameter) SetAttribute(Key string, Value interface{}) Node {
	n.attributes[Key] = Value
	return n
}

func (n Parameter) Position() *Position {
	return n.position
}

func (n Parameter) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Parameter) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VariableType != nil {
		vv := v.GetChildrenVisitor("VariableType")
		n.VariableType.Walk(vv)
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.DefaultValue != nil {
		vv := v.GetChildrenVisitor("DefaultValue")
		n.DefaultValue.Walk(vv)
	}

	v.LeaveNode(n)
}
