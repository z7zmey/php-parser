package node

type Parameter struct {
	name         string
	attributes   map[string]interface{}
	position     *Position
	variableType Node
	variable     Node
	defaultValue Node
}

func NewParameter(variableType Node, variable Node, defaultValue Node, byRef bool, variadic bool) Node {
	return Parameter{
		"Parameter",
		map[string]interface{}{
			"byRef":    byRef,
			"variadic": variadic,
		},
		nil,
		variableType,
		variable,
		defaultValue,
	}
}

func (n Parameter) Name() string {
	return "Parameter"
}

func (n Parameter) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Parameter) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Parameter) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Parameter) Position() *Position {
	return n.position
}

func (n Parameter) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Parameter) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

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
