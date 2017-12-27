package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Function) Name() string {
	return "Function"
}

type Function struct {
	name        string
	token       token.Token
	isReturnRef bool
	params      []node.Node
	returnType  node.Node
	stmts       []node.Node
}

func NewFunction(token token.Token, isReturnRef bool, params []node.Node, returnType node.Node, stmts []node.Node) node.Node {
	return Function{
		"Function",
		token,
		isReturnRef,
		params,
		returnType,
		stmts,
	}
}

func (n Function) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	v.Scalar("token", n.token.Value)
	v.Scalar("isReturnRef", n.isReturnRef)

	if n.params != nil {
		vv := v.Children("params")
		for _, nn := range n.params {
			nn.Walk(vv)
		}
	}

	if n.returnType != nil {
		vv := v.Children("returnType")
		n.returnType.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}
}
