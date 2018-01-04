package scalar

import (
	"github.com/z7zmey/php-parser/node"
)

type Encapsed struct {
	position *node.Position
	Parts    []node.Node
}

func NewEncapsed(Parts []node.Node) node.Node {
	return &Encapsed{
		nil,
		Parts,
	}
}

func (n Encapsed) Attributes() map[string]interface{} {
	return nil
}

func (n Encapsed) Position() *node.Position {
	return n.position
}

func (n Encapsed) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n Encapsed) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Parts != nil {
		vv := v.GetChildrenVisitor("Parts")
		for _, nn := range n.Parts {
			nn.Walk(vv)
		}
	}
}
