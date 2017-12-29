package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Static struct {
	name string
	vars []node.Node
}

func NewStatic(vars []node.Node) node.Node {
	return Static{
		"Static",
		vars,
	}
}

func (n Static) Name() string {
	return "Static"
}

func (n Static) Attributes() map[string]interface{} {
	return nil
}

func (n Static) Walk(v node.Visitor) {
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
