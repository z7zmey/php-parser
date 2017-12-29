package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Namespace struct {
	name          string
	namespaceName node.Node
	stmts         []node.Node
}

func NewNamespace(namespaceName node.Node, stmts []node.Node) node.Node {
	return Namespace{
		"Namespace",
		namespaceName,
		stmts,
	}
}

func (n Namespace) Name() string {
	return "Namespace"
}

func (n Namespace) Attributes() map[string]interface{} {
	return nil
}

func (n Namespace) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.namespaceName != nil {
		vv := v.GetChildrenVisitor("namespaceName")
		n.namespaceName.Walk(vv)
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
