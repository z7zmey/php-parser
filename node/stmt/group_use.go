package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n GroupUse) Name() string {
	return "GroupUse"
}

type GroupUse struct {
	name    string
	token   token.TokenInterface
	useType node.Node
	prefix  node.Node
	useList []node.Node
}

//TODO: stmts myst be []node.Node
func NewGroupUse(token token.TokenInterface, useType node.Node, prefix node.Node, useList []node.Node) node.Node {
	return GroupUse{
		"GroupUse",
		token,
		useType,
		prefix,
		useList,
	}
}

func (n GroupUse) SetToken(token token.TokenInterface) node.Node {
	n.token = token
	return n
}

func (n GroupUse) SetUseType(useType node.Node) node.Node {
	n.useType = useType
	return n
}

func (n GroupUse) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.useType != nil {
		vv := v.Children("useType")
		n.useType.Walk(vv)
	}

	if n.prefix != nil {
		vv := v.Children("prefix")
		n.prefix.Walk(vv)
	}

	if n.useList != nil {
		vv := v.Children("useList")
		for _, nn := range n.useList {
			nn.Walk(vv)
		}
	}
}
