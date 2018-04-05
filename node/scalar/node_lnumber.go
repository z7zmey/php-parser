package scalar

import "github.com/z7zmey/php-parser/walker"

// Lnumber node
type Lnumber struct {
	Value string
}

// NewLnumber node constructor
func NewLnumber(Value string) *Lnumber {
	return &Lnumber{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *Lnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Lnumber) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
