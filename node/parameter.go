package node

type Parameter struct {
	position     *Position
	ByRef        bool
	Variadic     bool
	VariableType Node
	Variable     Node
	DefaultValue Node
}

func NewParameter(VariableType Node, Variable Node, DefaultValue Node, ByRef bool, Variadic bool) *Parameter {
	return &Parameter{
		nil,
		ByRef,
		Variadic,
		VariableType,
		Variable,
		DefaultValue,
	}
}

func (n *Parameter) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef":    n.ByRef,
		"Variadic": n.Variadic,
	}
}

func (n *Parameter) Position() *Position {
	return n.position
}

func (n *Parameter) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n *Parameter) Walk(v Visitor) {
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
