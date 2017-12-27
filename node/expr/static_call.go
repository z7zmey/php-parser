package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n StaticCall) Name() string {
	return "StaticCall"
}

type StaticCall struct {
	name      string
	class     node.Node
	call      node.Node
	arguments []node.Node
}

func NewStaticCall(class node.Node, call node.Node, arguments []node.Node) node.Node {
	return StaticCall{
		"StaticCall",
		class,
		call,
		arguments,
	}
}

func (n StaticCall) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.class != nil {
		vv := v.Children("class")
		n.class.Walk(vv)
	}

	if n.call != nil {
		vv := v.Children("call")
		n.call.Walk(vv)
	}

	if n.arguments != nil {
		vv := v.Children("arguments")
		for _, nn := range n.arguments {
			nn.Walk(vv)
		}
	}
}
