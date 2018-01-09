package node

type Nullable struct {
	Expr Node
}

func NewNullable(Expression Node) *Nullable {
	return &Nullable{
		Expression,
	}
}

func (n *Nullable) Attributes() map[string]interface{} {
	return nil
}

func (n *Nullable) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
