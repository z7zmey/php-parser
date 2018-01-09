package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Namespace struct {
	NamespaceName node.Node
	Stmts         []node.Node
}

func NewNamespace(NamespaceName node.Node, Stmts []node.Node) *Namespace {
	return &Namespace{
		NamespaceName,
		Stmts,
	}
}

func (n *Namespace) Attributes() map[string]interface{} {
	return nil
}

func (n *Namespace) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.NamespaceName != nil {
		vv := v.GetChildrenVisitor("NamespaceName")
		n.NamespaceName.Walk(vv)
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
