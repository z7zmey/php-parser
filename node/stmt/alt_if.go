package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n AltIf) Name() string {
	return "AltIf"
}

type AltIf struct {
	name   string
	token  token.Token
	cond   node.Node
	stmt   node.Node
	elseIf []node.Node
	_else  node.Node
}

func NewAltIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltIf{
		"AltIf",
		token,
		cond,
		stmt,
		nil,
		nil,
	}
}

func (n AltIf) AddElseIf(elseIf node.Node) node.Node {
	if n.elseIf == nil {
		n.elseIf = make([]node.Node, 0)
	}

	n.elseIf = append(n.elseIf, elseIf)

	return n
}

func (n AltIf) SetElse(_else node.Node) node.Node {
	n._else = _else

	return n
}

func (n AltIf) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.GetChildrenVisitor("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.GetChildrenVisitor("stmt")
		n.stmt.Walk(vv)
	}

	if n.elseIf != nil {
		vv := v.GetChildrenVisitor("elseIf")
		for _, nn := range n.elseIf {
			nn.Walk(vv)
		}
	}

	if n._else != nil {
		vv := v.GetChildrenVisitor("else")
		n._else.Walk(vv)
	}

	v.LeaveNode(n)
}
