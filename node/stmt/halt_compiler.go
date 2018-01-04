package stmt

import (
	"github.com/z7zmey/php-parser/node"
)

type HaltCompiler struct {
	position *node.Position
}

func NewHaltCompiler() node.Node {
	return &HaltCompiler{
		nil,
	}
}

func (n HaltCompiler) Attributes() map[string]interface{} {
	return nil
}

func (n HaltCompiler) Position() *node.Position {
	return n.position
}

func (n HaltCompiler) SetPosition(p *node.Position) node.Node {
	n.position = p
	return n
}

func (n HaltCompiler) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	v.LeaveNode(n)
}
