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
		v.EnterChildNode("Condition", n)
		n.Condition.Walk(v)
		v.LeaveChildNode("Condition", n)
	}

	if n.IfTrue != nil {
		v.EnterChildNode("IfTrue", n)
		n.IfTrue.Walk(v)
		v.LeaveChildNode("IfTrue", n)
	}

	if n.IfFalse != nil {
		v.EnterChildNode("IfFalse", n)
		n.IfFalse.Walk(v)
		v.LeaveChildNode("IfFalse", n)
	}

	v.LeaveNode(n)
}
