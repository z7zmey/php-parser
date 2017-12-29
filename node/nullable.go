package node

type Nullable struct {
	name       string
	attributes map[string]interface{}
	expr       Node
}

func NewNullable(expression Node) Node {
	return Nullable{
		"Nullable",
		map[string]interface{}{},
		expression,
	}
}

func (n Nullable) Name() string {
	return "Nullable"
}

func (n Nullable) Attributes() map[string]interface{} {
	return n.attributes
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
