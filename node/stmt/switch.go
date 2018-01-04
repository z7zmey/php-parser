package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Switch struct {
	attributes map[string]interface{}
	position   *node.Position
	token      token.Token
	cond       node.Node
	cases      []node.Node
}

func NewSwitch(token token.Token, cond node.Node, cases []node.Node) node.Node {
	return &Switch{
		map[string]interface{}{},
		nil,
		token,
		cond,
		cases,
	}
}

func (n Switch) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Switch) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Switch) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
}

func (n Switch) Position() *node.Position {
	return n.position
}

func (n Switch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
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
