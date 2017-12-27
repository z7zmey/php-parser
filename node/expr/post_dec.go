package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PostDec) Name() string {
	return "PostDec"
}

type PostDec struct {
	name     string
	variable node.Node
}

func NewPostDec(variableession node.Node) node.Node {
	return PostDec{
		"PostDec",
		variableession,
	}
}

func (n PostDec) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.Children("variable")
		n.variable.Walk(vv)
	}
}
