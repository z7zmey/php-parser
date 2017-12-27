package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Namespace) Name() string {
	return "Namespace"
}

type Namespace struct {
	name          string
	token         token.Token
	namespaceName node.Node
	stmts         []node.Node
}

func NewNamespace(token token.Token, namespaceName node.Node, stmts []node.Node) node.Node {
	return Namespace{
		"Namespace",
		token,
		namespaceName,
		stmts,
	}
}

func (n Namespace) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.namespaceName != nil {
		vv := v.Children("namespaceName")
		n.namespaceName.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
