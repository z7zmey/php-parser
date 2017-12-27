package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n PostInc) Name() string {
	return "PostInc"
}

type PostInc struct {
	name     string
	variable node.Node
}

func NewPostInc(variableession node.Node) node.Node {
	return PostInc{
		"PostInc",
		variableession,
	}
}

func (n PostInc) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variable != nil {
		vv := v.Children("variable")
		n.variable.Walk(vv)
	}
}
