package stmt

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type Switch struct {
	position *node.Position
	comments *[]comment.Comment
	token    token.Token
	Cond     node.Node
	cases    []node.Node
}

func NewSwitch(token token.Token, Cond node.Node, cases []node.Node) *Switch {
	return &Switch{
		nil,
		nil,
		token,
		Cond,
		cases,
	}
}

func (n Switch) Attributes() map[string]interface{} {
	return nil
}

func (n Switch) Position() *node.Position {
	return n.position
}

func (n Switch) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Switch) Comments() *[]comment.Comment {
	return n.comments
}

func (n Switch) SetComments(c *[]comment.Comment) node.Node {
	n.comments = c
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
