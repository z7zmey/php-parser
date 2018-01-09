package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Label struct {
	LabelName node.Node
}

func NewLabel(LabelName node.Node) *Label {
	return &Label{
		LabelName,
	}
}

func (n *Label) Attributes() map[string]interface{} {
	return nil
}

func (n *Label) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.LabelName != nil {
		vv := v.GetChildrenVisitor("LabelName")
		n.LabelName.Walk(vv)
	}

	v.LeaveNode(n)
}
