package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type Catch struct {
	Types    []node.Node
	Variable node.Node
	Stmts    []node.Node
}

func NewCatch(Types []node.Node, Variable node.Node, Stmts []node.Node) *Catch {
	return &Catch{
		Types,
		Variable,
		Stmts,
	}
}

func (n *Catch) Attributes() map[string]interface{} {
	return nil
}

func (n *Catch) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Types != nil {
		vv := v.GetChildrenVisitor("Types")
		for _, nn := range n.Types {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	if n.Stmts != nil {
		vv := v.GetChildrenVisitor("Stmts")
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(vv)
			}
		}
	}

	v.LeaveNode(n)
}
