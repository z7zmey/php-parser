package node

type Argument struct {
	name      string
	arguments map[string]interface{}
	expr      Node
	variadic  bool
}

func NewArgument(expression Node, variadic bool) Node {
	return Argument{
		"Argument",
		map[string]interface{}{
			"variadic": variadic,
		},
		expression,
		variadic,
	}
}

func (n Argument) Name() string {
	return "Argument"
}

func (n Argument) Attributes() map[string]interface{} {
	return n.arguments
}

func (n Argument) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}
