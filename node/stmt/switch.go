package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Switch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Switch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Switch) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Switch) Name() string {
	return "Switch"
}

type Switch struct {
	name       string
	attributes map[string]interface{}
	token      token.Token
	cond       node.Node
	cases      []node.Node
}

func NewSwitch(token token.Token, cond node.Node, cases []node.Node) node.Node {
	return Switch{
		"Switch",
		map[string]interface{}{},
		token,
		cond,
		cases,
	}
}

func (n Switch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.cases != nil {
		vv := v.GetChildrenVisitor("cases")
		for _, nn := range n.cases {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
