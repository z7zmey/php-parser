package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Goto struct {
	name       string
	attributes map[string]interface{}
	label      node.Node
}

func NewGoto(label node.Node) node.Node {
	return Goto{
		"Goto",
		map[string]interface{}{},
		label,
	}
}

func (n Goto) Name() string {
	return "Goto"
}

func (n Goto) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Goto) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Goto) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Goto) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.label != nil {
		vv := v.GetChildrenVisitor("label")
		n.label.Walk(vv)
	}

	v.LeaveNode(n)
}
