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
		vv := v.GetChildrenVisitor("ClassName")
		n.ClassName.Walk(vv)
	}

	if n.Modifiers != nil {
		vv := v.GetChildrenVisitor("Modifiers")
		for _, nn := range n.Modifiers {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.ArgumentList != nil {
		vv := v.GetChildrenVisitor("ArgumentList")
		n.ArgumentList.Walk(vv)
	}

	if n.Extends != nil {
		vv := v.GetChildrenVisitor("Extends")
		n.Extends.Walk(vv)
	}

	if n.Implements != nil {
		vv := v.GetChildrenVisitor("Implements")
		n.Implements.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
