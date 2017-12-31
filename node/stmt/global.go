package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	name       string
	attributes map[string]interface{}
	vars       []node.Node
}

func NewGlobal(vars []node.Node) node.Node {
	return Global{
		"Global",
		map[string]interface{}{},
		vars,
	}
}

func (n Global) Name() string {
	return "Global"
}

func (n Global) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Global) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Global) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Global) Walk(v node.Visitor) {
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
