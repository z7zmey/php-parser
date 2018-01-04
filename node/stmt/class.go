package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Class struct {
	attributes map[string]interface{}
	position   *node.Position
	ClassName  node.Node
	Modifiers  []node.Node
	args       []node.Node
	Extends    node.Node
	Implements []node.Node
	Stmts      []node.Node
}

func NewClass(ClassName node.Node, Modifiers []node.Node, args []node.Node, Extends node.Node, Implements []node.Node, Stmts []node.Node, phpDocComment string) node.Node {
	return &Class{
		map[string]interface{}{
			"phpDocComment": phpDocComment,
		},
		nil,
		ClassName,
		Modifiers,
		args,
		Extends,
		Implements,
		Stmts,
	}
}

func (n Class) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Class) Position() *node.Position {
	return n.position
}

func (n Class) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Class) Walk(v node.Visitor) {
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
			nn.Walk(vv)
		}
	}

	if n.args != nil {
		vv := v.GetChildrenVisitor("args")
		for _, nn := range n.args {
			nn.Walk(vv)
		}
	}

	if n.Extends != nil {
		vv := v.GetChildrenVisitor("Extends")
		n.Extends.Walk(vv)
	}

	if n.Implements != nil {
		vv := v.GetChildrenVisitor("Implements")
		for _, nn := range n.Implements {
			nn.Walk(vv)
		}
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
