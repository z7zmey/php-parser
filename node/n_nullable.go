package node

import "github.com/z7zmey/php-parser/walker"

// Nullable node
type Nullable struct {
	Expr Node
}

// NewNullable node constructor
func NewNullable(Expression Node) *Nullable {
	return &Nullable{
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Nullable) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Nullable) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
