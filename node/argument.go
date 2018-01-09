package node

type Argument struct {
	Variadic bool
	Expr     Node
}

func NewArgument(Expression Node, Variadic bool) *Argument {
	return &Argument{
		Variadic,
		Expression,
	}
}

func (n *Argument) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Variadic": n.Variadic,
	}
}

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
