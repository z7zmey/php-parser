package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Finally struct {
	name       string
	attributes map[string]interface{}
	stmts      []node.Node
}

func NewFinally(stmts []node.Node) node.Node {
	return Finally{
		"Finally",
		map[string]interface{}{},
		stmts,
	}
}

func (n Finally) Name() string {
	return "Finally"
}

func (n Finally) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Finally) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
