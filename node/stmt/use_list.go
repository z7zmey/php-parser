package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n UseList) Name() string {
	return "UseList"
}

type UseList struct {
	name    string
	token   token.Token
	useType node.Node
	uses    []node.Node
}

func NewUseList(token token.Token, useType node.Node, uses []node.Node) node.Node {
	return UseList{
		"UseList",
		token,
		useType,
		uses,
	}
}

func (n UseList) Walk(v node.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.GetChildrenVisitor("useType")
		n.useType.Walk(vv)
	}

	if n.uses != nil {
		vv := v.GetChildrenVisitor("uses")
		for _, nn := range n.uses {
			nn.Walk(vv)
		}
	}

	v.LeaveNode(n)
}
