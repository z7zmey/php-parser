package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type AltIf struct {
	If
}

func NewAltIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltIf{
		If{
			node.SimpleNode{Name: "AltIf", Attributes: make(map[string]string)},
			token,
			cond,
			stmt,
			nil,
			nil,
		},
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
