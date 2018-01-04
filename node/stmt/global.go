package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Global struct {
	position *node.Position
	Vars     []node.Node
}

func NewGlobal(Vars []node.Node) node.Node {
	return &Global{
		nil,
		Vars,
	}
}

func (n Global) Attributes() map[string]interface{} {
	return nil
}

func (n Global) Position() *node.Position {
	return n.position
}

func (n Global) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Global) Walk(v node.Visitor) {
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
