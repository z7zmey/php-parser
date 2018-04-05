package scalar

import "github.com/z7zmey/php-parser/walker"

// Dnumber node
type Dnumber struct {
	Value string
}

// NewDnumber node constructor
func NewDnumber(Value string) *Dnumber {
	return &Dnumber{
		Value,
	}
}

// Attributes returns node attributes as map
func (n *Dnumber) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"Value": n.Value,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Dnumber) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
