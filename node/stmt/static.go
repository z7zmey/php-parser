package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Static struct {
	name       string
	attributes map[string]interface{}
	vars       []node.Node
}

func NewStatic(vars []node.Node) node.Node {
	return Static{
		"Static",
		map[string]interface{}{},
		vars,
	}
}

func (n Static) Name() string {
	return "Static"
}

func (n Static) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Static) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Static) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Static) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.GetChildrenVisitor("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
