package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Label struct {
	name      string
	labelName node.Node
}

func NewLabel(labelName node.Node) node.Node {
	return Label{
		"Label",
		labelName,
	}
}

func (n Label) Name() string {
	return "Label"
}

func (n Label) Attributes() map[string]interface{} {
	return nil
}

func (n Label) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.labelName != nil {
		vv := v.GetChildrenVisitor("labelName")
		n.labelName.Walk(vv)
	}

	v.LeaveNode(n)
}
