package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type InlineHtml struct {
	name       string
	attributes map[string]interface{}
}

func NewInlineHtml(value string) node.Node {
	return InlineHtml{
		"InlineHtml",
		map[string]interface{}{
			"value": value,
		},
	}
}

func (n InlineHtml) Name() string {
	return "InlineHtml"
}

func (n InlineHtml) Attributes() map[string]interface{} {
	return n.attributes
}

func (n InlineHtml) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n InlineHtml) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n InlineHtml) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
