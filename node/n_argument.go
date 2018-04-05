package node

import "github.com/z7zmey/php-parser/walker"

// Argument node
type Argument struct {
	Variadic    bool // if ... before variable
	IsReference bool // if & before variable
	Expr        Node // Exression
}

// NewArgument node constructor
func NewArgument(Expression Node, Variadic bool, IsReference bool) *Argument {
	return &Argument{
		Variadic,
		IsReference,
		Expression,
	}
}

// Attributes returns node attributes as map
func (n *Argument) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Variadic":    n.Variadic,
		"IsReference": n.IsReference,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Argument) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Expr != nil {
		vv := v.GetChildrenVisitor("Expr")
		n.Expr.Walk(vv)
	}

	v.LeaveNode(n)
}
