package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Constant node
type Constant struct {
	PhpDocComment string
	ConstantName  node.Node
	Expr          node.Node
}

// NewConstant node constructor
func NewConstant(ConstantName node.Node, Expr node.Node, PhpDocComment string) *Constant {
	return &Constant{
		PhpDocComment,
		ConstantName,
		Expr,
	}
}

// Attributes returns node attributes as map
func (n *Constant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Constant) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		vv := v.GetChildrenVisitor("ConstantName")
		n.ConstantName.Walk(vv)
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
