package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Goto struct {
	Label node.Node
}

func NewGoto(Label node.Node) *Goto {
	return &Goto{
		Label,
	}
}

func (n *Goto) Attributes() map[string]interface{} {
	return nil
}

func (n *Goto) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Label != nil {
		vv := v.GetChildrenVisitor("Label")
		n.Label.Walk(vv)
	}

	v.LeaveNode(n)
}
