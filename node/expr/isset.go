package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Isset) Name() string {
	return "Isset"
}

type Isset struct {
	name      string
	variables []node.Node
}

func NewIsset(variables []node.Node) node.Node {
	return Isset{
		"Isset",
		variables,
	}
}

func (n Isset) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.variables != nil {
		vv := v.Children("variables")
		for _, nn := range n.variables {
			nn.Walk(vv)
		}
	}
}
