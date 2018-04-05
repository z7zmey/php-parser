package stmt

import "github.com/z7zmey/php-parser/walker"

// InlineHtml node
type InlineHtml struct {
	Value string
}

// NewInlineHtml node constructor
func NewInlineHtml(Value string) *InlineHtml {
	return &InlineHtml{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *InlineHtml) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InlineHtml) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
