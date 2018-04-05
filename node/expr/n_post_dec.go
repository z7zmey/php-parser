package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PostDec node
type PostDec struct {
	Variable node.Node
}

// NewPostDec node constructor
func NewPostDec(Variable node.Node) *PostDec {
	return &PostDec{
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *PostDec) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PostDec) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
