package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n For) Name() string {
	return "For"
}

type For struct {
	name  string
	token token.Token
	init  []node.Node
	cond  []node.Node
	loop  []node.Node
	stmt  node.Node
}

func NewFor(token token.Token, init []node.Node, cond []node.Node, loop []node.Node, stmt node.Node) node.Node {
	return For{
		"For",
		token,
		init,
		cond,
		loop,
		stmt,
	}
}

func (n For) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.init != nil {
		vv := v.GetChildrenVisitor("init")
		for _, nn := range n.init {
			nn.Walk(vv)
		}
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		for _, nn := range n.cond {
			nn.Walk(vv)
		}
	}

	if n.loop != nil {
		vv := v.GetChildrenVisitor("loop")
		for _, nn := range n.loop {
			nn.Walk(vv)
		}
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	v.LeaveNode(n)
}
