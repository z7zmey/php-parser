package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Foreach node
type Foreach struct {
	Expr     node.Node
	Key      node.Node
	Variable node.Node
	Stmt     node.Node
}

// NewForeach node constructor
func NewForeach(Expr node.Node, Key node.Node, Variable node.Node, Stmt node.Node) *Foreach {
	return &Foreach{
		Expr,
		Key,
		Variable,
		Stmt,
	}
}

// Attributes returns node attributes as map
func (n *Foreach) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Foreach) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	if n.Key != nil {
		vv := v.GetChildrenVisitor("Key")
		n.Key.Walk(vv)
	}

	if n.Variable != nil {
		vv := v.GetChildrenVisitor("Variable")
		n.Variable.Walk(vv)
	}

	if n.Stmt != nil {
		vv := v.GetChildrenVisitor("Stmt")
		n.Stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
