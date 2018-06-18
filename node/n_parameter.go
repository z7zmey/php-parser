package node

import "github.com/z7zmey/php-parser/walker"

// Parameter node
type Parameter struct {
	ByRef        bool
	Variadic     bool
	VariableType Node
	Variable     Node
	DefaultValue Node
}

// NewParameter node constructor
func NewParameter(VariableType Node, Variable Node, DefaultValue Node, ByRef bool, Variadic bool) *Parameter {
	return &Parameter{
		ByRef,
		Variadic,
		VariableType,
		Variable,
		DefaultValue,
	}
}

// Attributes returns node attributes as map
func (n *Parameter) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ByRef":    n.ByRef,
		"Variadic": n.Variadic,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Parameter) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.VariableType != nil {
		v.EnterChildNode("VariableType", n)
		n.VariableType.Walk(v)
		v.LeaveChildNode("VariableType", n)
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.DefaultValue != nil {
		v.EnterChildNode("DefaultValue", n)
		n.DefaultValue.Walk(v)
		v.LeaveChildNode("DefaultValue", n)
	}

	v.LeaveNode(n)
}
