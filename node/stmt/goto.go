package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Goto struct {
	name  string
	label node.Node
}

func NewGoto(label node.Node) node.Node {
	return Goto{
		"Goto",
		label,
	}
}

func (n Goto) Name() string {
	return "Goto"
}

func (n Goto) Attributes() map[string]interface{} {
	return nil
}

func (n Goto) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.label != nil {
		vv := v.GetChildrenVisitor("label")
		n.label.Walk(vv)
	}

	v.LeaveNode(n)
}
