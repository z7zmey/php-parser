package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n ConstFetch) Name() string {
	return "ConstFetch"
}

type ConstFetch struct {
	name     string
	constant node.Node
}

func NewConstFetch(constant node.Node) node.Node {
	return ConstFetch{
		"ConstFetch",
		constant,
	}
}

func (n ConstFetch) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.constant != nil {
		vv := v.Children("constant")
		n.constant.Walk(vv)
	}
}
