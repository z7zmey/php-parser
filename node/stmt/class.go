package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Class struct {
	name       string
	attributes map[string]interface{}
	className  node.Node
	modifiers  []node.Node
	args       []node.Node
	extends    node.Node
	implements []node.Node
	stmts      []node.Node
}

func NewClass(className node.Node, modifiers []node.Node, args []node.Node, extends node.Node, implements []node.Node, stmts []node.Node) node.Node {
	return Class{
		"Class",
		map[string]interface{}{},
		className,
		modifiers,
		args,
		extends,
		implements,
		stmts,
	}
}

func (n Class) Name() string {
	return "Class"
}

func (n Class) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Class) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Class) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Class) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.className != nil {
		vv := v.GetChildrenVisitor("className")
		n.className.Walk(vv)
	}

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.args != nil {
		vv := v.GetChildrenVisitor("args")
		for _, nn := range n.args {
			nn.Walk(vv)
		}
	}

	if n.extends != nil {
		vv := v.GetChildrenVisitor("extends")
		n.extends.Walk(vv)
	}

	if n.implements != nil {
		vv := v.GetChildrenVisitor("implements")
		for _, nn := range n.implements {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
