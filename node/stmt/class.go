package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Class) Name() string {
	return "Class"
}

type Class struct {
	name       string
	token      token.Token
	modifiers  []node.Node
	args       []node.Node
	extends    node.Node
	implements []node.Node
	stmts      []node.Node
}

func NewClass(token token.Token, modifiers []node.Node, args []node.Node, extends node.Node, implements []node.Node, stmts []node.Node) node.Node {
	return Class{
		"Class",
		token,
		modifiers,
		args,
		extends,
		implements,
		stmts,
	}
}

func (n Class) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)

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
