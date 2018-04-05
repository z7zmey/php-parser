package stmt

import "github.com/z7zmey/php-parser/walker"

// HaltCompiler node
type HaltCompiler struct {
}

// NewHaltCompiler node constructor
func NewHaltCompiler() *HaltCompiler {
	return &HaltCompiler{}
}

// Attributes returns node attributes as map
func (n *HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *HaltCompiler) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
