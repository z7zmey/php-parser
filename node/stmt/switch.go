package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Switch struct {
	attributes map[string]interface{}
	position   *node.Position
	token      token.Token
	Cond       node.Node
	cases      []node.Node
}

func NewSwitch(token token.Token, Cond node.Node, cases []node.Node) node.Node {
	return &Switch{
		map[string]interface{}{},
		nil,
		token,
		Cond,
		cases,
	}
}

func (n Switch) Attributes() map[string]interface{} {
	return n.attributes
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

	if n.Cond != nil {
		vv := v.GetChildrenVisitor("Cond")
		n.Cond.Walk(vv)
	}

	if n.cases != nil {
		vv := v.GetChildrenVisitor("cases")
		for _, nn := range n.cases {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
