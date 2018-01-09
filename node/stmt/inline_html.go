package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type InlineHtml struct {
	Value string
}

func NewInlineHtml(Value string) *InlineHtml {
	return &InlineHtml{
		Value,
	}
}

func (n *InlineHtml) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

func (n *InlineHtml) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
