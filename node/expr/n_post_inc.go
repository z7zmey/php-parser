package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// PostInc node
type PostInc struct {
	Variable node.Node
}

// NewPostInc node constructor
func NewPostInc(Variable node.Node) *PostInc {
	return &PostInc{
		Variable,
	}
}

// Attributes returns node attributes as map
func (n *PostInc) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *PostInc) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	v.LeaveNode(n)
}
