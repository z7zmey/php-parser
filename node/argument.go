package node

// Argument node
type Argument struct {
	Variadic bool // if ... before variable
	Expr     Node // Exression
}

// NewArgument node constuctor
func NewArgument(Expression Node, Variadic bool) *Argument {
	return &Argument{
		Variadic,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Argument) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Variadic": n.Variadic,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Argument) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
