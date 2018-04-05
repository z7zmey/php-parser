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
