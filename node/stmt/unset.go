package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Unset) Name() string {
	return "Unset"
}

type Unset struct {
	name  string
	token token.Token
	vars  []node.Node
}

func NewUnset(token token.Token, vars []node.Node) node.Node {
	return Unset{
		"Unset",
		token,
		vars,
	}
}

func (n Unset) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.vars != nil {
		vv := v.GetChildrenVisitor("vars")
		for _, nn := range n.vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
