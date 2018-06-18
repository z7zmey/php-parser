package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/walker"
)

// Class node
type Class struct {
	PhpDocComment string
	ClassName     node.Node
	Modifiers     []node.Node
	ArgumentList  *node.ArgumentList
	Extends       *ClassExtends
	Implements    *ClassImplements
	Stmts         []node.Node
}

// NewClass node constructor
func NewClass(ClassName node.Node, Modifiers []node.Node, ArgumentList *node.ArgumentList, Extends *ClassExtends, Implements *ClassImplements, Stmts []node.Node, PhpDocComment string) *Class {
	return &Class{
		PhpDocComment,
		ClassName,
		Modifiers,
		ArgumentList,
		Extends,
		Implements,
		Stmts,
	}
}

// Attributes returns node attributes as map
func (n *Class) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Class) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ClassName != nil {
		v.EnterChildNode("ClassName", n)
		n.ClassName.Walk(v)
		v.LeaveChildNode("ClassName", n)
	}

	if n.Modifiers != nil {
		v.EnterChildList("Modifiers", n)
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Modifiers", n)
	}

	if n.ArgumentList != nil {
		v.EnterChildNode("ArgumentList", n)
		n.ArgumentList.Walk(v)
		v.LeaveChildNode("ArgumentList", n)
	}

	if n.Extends != nil {
		v.EnterChildNode("Extends", n)
		n.Extends.Walk(v)
		v.LeaveChildNode("Extends", n)
	}

	if n.Implements != nil {
		v.EnterChildNode("Implements", n)
		n.Implements.Walk(v)
		v.LeaveChildNode("Implements", n)
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
