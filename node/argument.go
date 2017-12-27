package node

type Argument struct {
	name     string
	expr     Node
	variadic bool
}

func (n Argument) Name() string {
	return "Argument"
}

func NewArgument(expression Node, variadic bool) Node {
	return Argument{
		"Argument",
		expression,
		variadic,
	}
}

func (n Argument) Walk(v Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
