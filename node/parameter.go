package node

type Parameter struct {
	name         string
	variableType Node
	variable     Node
	defaultValue Node
	byRef        bool
	variadic     bool
}

func (n Parameter) Name() string {
	return "Parameter"
}

func NewParameter(variableType Node, variable Node, defaultValue Node, byRef bool, variadic bool) Node {
	return Parameter{
		"Parameter",
		variableType,
		variable,
		defaultValue,
		byRef,
		variadic,
	}
}

func (n Parameter) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("byRef", n.byRef)
	v.Scalar("variadic", n.variadic)

	if n.variableType != nil {
		vv := v.GetChildrenVisitor("variableType")
		n.variableType.Walk(vv)
	}

	if n.variable != nil {
		vv := v.GetChildrenVisitor("variable")
		n.variable.Walk(vv)
	}

	if n.defaultValue != nil {
		vv := v.GetChildrenVisitor("defaultValue")
		n.defaultValue.Walk(vv)
	}

	v.LeaveNode(n)
}
