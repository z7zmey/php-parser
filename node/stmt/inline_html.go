package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type InlineHtml struct {
	attributes map[string]interface{}
	position   *node.Position
}

func NewInlineHtml(value string) node.Node {
	return &InlineHtml{
		map[string]interface{}{
			"value": value,
		},
		nil,
	}
}

func (n InlineHtml) Attributes() map[string]interface{} {
	return n.attributes
}

func (n InlineHtml) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n InlineHtml) SetAttribute(key string, value interface{}) node.Node {
	n.attributes[key] = value
	return n
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
