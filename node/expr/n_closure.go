package expr

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Closure node
type Closure struct {
	ReturnsRef    bool
	Static        bool
	PhpDocComment string
	Params        []node.Node
	ClosureUse    *ClosureUse
	ReturnType    node.Node
	Stmts         []node.Node
}

// NewClosure node constructor
func NewClosure(Params []node.Node, ClosureUse *ClosureUse, ReturnType node.Node, Stmts []node.Node, Static bool, ReturnsRef bool, PhpDocComment string) *Closure {
	return &Closure{
		ReturnsRef,
		Static,
		PhpDocComment,
		Params,
		ClosureUse,
		ReturnType,
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *Closure) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"ReturnsRef":    n.ReturnsRef,
		"Static":        n.Static,
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Closure) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Params != nil {
		v.EnterChildList("Params", n)
		for _, nn := range n.Params {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Params", n)
	}

	if n.ClosureUse != nil {
		v.EnterChildNode("ClosureUse", n)
		n.ClosureUse.Walk(v)
		v.LeaveChildNode("ClosureUse", n)
	}

	if n.ReturnType != nil {
		v.EnterChildNode("ReturnType", n)
		n.ReturnType.Walk(v)
		v.LeaveChildNode("ReturnType", n)
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}
