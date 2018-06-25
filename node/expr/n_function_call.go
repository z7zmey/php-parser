package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// FunctionCall node
type FunctionCall struct {
	Comments     []*comment.Comment
	Position     *position.Position
	Function     node.Node
	ArgumentList *node.ArgumentList
}

// NewFunctionCall node constructor
func NewFunctionCall(Function node.Node, ArgumentList *node.ArgumentList) *FunctionCall {
	return &FunctionCall{
		Function:     Function,
		ArgumentList: ArgumentList,
	}
}

// SetPosition sets node position
func (n *FunctionCall) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *FunctionCall) GetPosition() *position.Position {
	return n.Position
}

func (n *FunctionCall) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *FunctionCall) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *FunctionCall) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *FunctionCall) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Function != nil {
		v.EnterChildNode("Function", n)
		n.Function.Walk(v)
		v.LeaveChildNode("Function", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	v.LeaveNode(n)
}
