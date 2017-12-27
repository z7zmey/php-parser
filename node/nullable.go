package node

type Nullable struct {
	name string
	expr Node
}

func (n Nullable) Name() string {
	return "Nullable"
}

func NewNullable(expression Node) Node {
	return Nullable{
		"Nullable",
		expression,
	}
}

func (n Nullable) Walk(v Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
