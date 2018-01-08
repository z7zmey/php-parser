package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
)

type FunctionCall struct {
	position  *node.Position
	comments  []comment.Comment
	Function  node.Node
	Arguments []node.Node
}

func NewFunctionCall(Function node.Node, Arguments []node.Node) *FunctionCall {
	return &FunctionCall{
		nil,
		nil,
		Function,
		Arguments,
	}
}

func (n *FunctionCall) Attributes() map[string]interface{} {
	return nil
}

func (n *FunctionCall) Position() *node.Position {
	return n.position
}

func (n *FunctionCall) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n *FunctionCall) Comments() []comment.Comment {
	return n.comments
}

func (n *FunctionCall) SetComments(c []comment.Comment) node.Node {
	n.comments = c
	return n
}

func (n *FunctionCall) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Function != nil {
		vv := v.GetChildrenVisitor("Function")
		n.Function.Walk(vv)
	}

	if n.Arguments != nil {
		vv := v.GetChildrenVisitor("Arguments")
		for _, nn := range n.Arguments {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
