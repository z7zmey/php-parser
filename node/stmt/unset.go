package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Unset struct {
	name string
	vars []node.Node
}

func NewUnset(vars []node.Node) node.Node {
	return Unset{
		"Unset",
		vars,
	}
}

func (n Unset) Name() string {
	return "Unset"
}

func (n Unset) Attributes() map[string]interface{} {
	return nil
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
