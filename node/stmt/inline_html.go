package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type InlineHtml struct {
	position *node.Position
	Value    string
}

func NewInlineHtml(Value string) node.Node {
	return &InlineHtml{
		nil,
		Value,
	}
}

func (n InlineHtml) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n InlineHtml) Position() *node.Position {
	return n.position
}

func (n InlineHtml) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n InlineHtml) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
