package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Catch struct {
	name     string
	types    []node.Node
	variable node.Node
	stmts    []node.Node
}

func NewCatch(types []node.Node, variable node.Node, stmts []node.Node) node.Node {
	return Catch{
		"Catch",
		types,
		variable,
		stmts,
	}
}

func (n Catch) Name() string {
	return "Catch"
}

func (n Catch) Attributes() map[string]interface{} {
	return nil
}

func (n Catch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.types != nil {
		vv := v.GetChildrenVisitor("types")
		for _, nn := range n.types {
			nn.Walk(vv)
		}
	}

	if n.stmts != nil {
		vv := v.GetChildrenVisitor("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
