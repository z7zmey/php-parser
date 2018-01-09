package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	Vars []node.Node
}

func NewGlobal(Vars []node.Node) *Global {
	return &Global{
		Vars,
	}
}

func (n *Global) Attributes() map[string]interface{} {
	return nil
}

func (n *Global) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Vars != nil {
		vv := v.GetChildrenVisitor("Vars")
		for _, nn := range n.Vars {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
