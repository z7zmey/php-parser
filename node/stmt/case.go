package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Case struct {
	name       string
	attributes map[string]interface{}
	cond       node.Node
	stmts      []node.Node
}

func NewCase(cond node.Node, stmts []node.Node) node.Node {
	return Case{
		"Case",
		map[string]interface{}{},
		cond,
		stmts,
	}
}

func (n Case) Name() string {
	return "Case"
}

func (n Case) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Case) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Case) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Case) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
