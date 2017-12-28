package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n ClassConstList) Name() string {
	return "ClassConstList"
}

type ClassConstList struct {
	name      string
	token     token.Token
	modifiers []node.Node
	consts    []node.Node
}

func NewClassConstList(token token.Token, modifiers []node.Node, consts []node.Node) node.Node {
	return ClassConstList{
		"ClassConstList",
		token,
		modifiers,
		consts,
	}
}

func (n ClassConstList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.modifiers != nil {
		vv := v.GetChildrenVisitor("modifiers")
		for _, nn := range n.modifiers {
			nn.Walk(vv)
		}
	}

	if n.consts != nil {
		vv := v.GetChildrenVisitor("consts")
		for _, nn := range n.consts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
