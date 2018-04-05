package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Ternary node
type Ternary struct {
	Condition node.Node
	IfTrue    node.Node
	IfFalse   node.Node
}

// NewTernary node constructor
func NewTernary(Condition node.Node, IfTrue node.Node, IfFalse node.Node) *Ternary {
	return &Ternary{
		Condition,
		IfTrue,
		IfFalse,
	}
}

// Attributes returns node attributes as map
func (n *Ternary) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Ternary) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Condition != nil {
		vv := v.GetChildrenVisitor("Condition")
		n.Condition.Walk(vv)
	}

	if n.IfTrue != nil {
		vv := v.GetChildrenVisitor("IfTrue")
		n.IfTrue.Walk(vv)
	}

	if n.IfFalse != nil {
		vv := v.GetChildrenVisitor("IfFalse")
		n.IfFalse.Walk(vv)
	}

	v.LeaveNode(n)
}
