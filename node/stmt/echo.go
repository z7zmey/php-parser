package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Echo struct {
	name       string
	attributes map[string]interface{}
	exprs      []node.Node
}

func NewEcho(exprs []node.Node) node.Node {
	return Echo{
		"Echo",
		map[string]interface{}{},
		exprs,
	}
}

func (n Echo) Name() string {
	return "Echo"
}

func (n Echo) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Echo) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Echo) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Echo) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.exprs != nil {
		vv := v.GetChildrenVisitor("exprs")
		for _, nn := range n.exprs {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
