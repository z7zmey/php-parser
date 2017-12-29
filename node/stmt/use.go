package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Use struct {
	name    string
	useType node.Node
	use     node.Node
	alias   node.Node
}

func NewUse(useType node.Node, use node.Node, alias node.Node) node.Node {
	return Use{
		"Use",
		useType,
		use,
		alias,
	}
}

func (n Use) Name() string {
	return "Use"
}

func (n Use) Attributes() map[string]interface{} {
	return nil
}

func (n Use) SetUseType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n Use) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.GetChildrenVisitor("useType")
		n.useType.Walk(vv)
	}

	if n.use != nil {
		vv := v.GetChildrenVisitor("use")
		n.use.Walk(vv)
	}

	if n.alias != nil {
		vv := v.GetChildrenVisitor("alias")
		n.alias.Walk(vv)
	}

	v.LeaveNode(n)
}
