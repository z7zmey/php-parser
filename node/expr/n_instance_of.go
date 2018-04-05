package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// InstanceOf node
type InstanceOf struct {
	Expr  node.Node
	Class node.Node
}

// NewInstanceOf node constructor
func NewInstanceOf(Expr node.Node, Class node.Node) *InstanceOf {
	return &InstanceOf{
		Expr,
		Class,
	}
}

// Attributes returns node attributes as map
func (n *InstanceOf) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InstanceOf) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	if n.Class != nil {
		vv := v.GetChildrenVisitor("Class")
		n.Class.Walk(vv)
	}

	v.LeaveNode(n)
}
