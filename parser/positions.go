package parser

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func getListPosStartLine(l []node.Node) int {
	startLine := -1

	if l == nil {
		return startLine
	}

	if len(l) == 0 {
		return startLine
	}

	return getNodePosStartLine(l[0])
}

func getNodePosStartLine(n node.Node) int {
	startLine := -1

	if n == nil {
		return startLine
	}

	p := n.Position()
	if p != nil {
		startLine = p.StartLine
	}

	return startLine
}

func getListPosEndLine(l []node.Node) int {
	endLine := -1

	if l == nil {
		return endLine
	}

	if len(l) == 0 {
		return endLine
	}

	return getNodePosEndLine(l[len(l)-1])
}

func getNodePosEndLine(n node.Node) int {
	endLine := -1

	if n == nil {
		return endLine
	}

	p := n.Position()
	if p != nil {
		endLine = p.EndLine
	}

	return endLine
}

func NewNodeListPosition(list []node.Node) *node.Position {
	return &node.Position{getListPosStartLine(list), getListPosEndLine(list)}
}

func NewNodePosition(n node.Node) *node.Position {
	return &node.Position{getNodePosStartLine(n), getNodePosEndLine(n)}
}

func NewTokenPosition(t token.Token) *node.Position {
	return &node.Position{t.StartLine, t.EndLine}
}

func NewTokensPosition(startToken token.Token, EndToken token.Token) *node.Position {
	return &node.Position{startToken.StartLine, EndToken.EndLine}
}

func NewTokenNodePosition(t token.Token, n node.Node) *node.Position {
	return &node.Position{t.StartLine, getNodePosEndLine(n)}
}

func NewNodeTokenPosition(n node.Node, t token.Token) *node.Position {
	return &node.Position{getNodePosStartLine(n), t.EndLine}
}

func NewNodesPosition(startNode node.Node, endNode node.Node) *node.Position {
	return &node.Position{getNodePosStartLine(startNode), getNodePosEndLine(endNode)}
}

func NewNodeListTokenPosition(list []node.Node, t token.Token) *node.Position {
	return &node.Position{getListPosStartLine(list), t.EndLine}
}

func NewTokenNodeListPosition(t token.Token, list []node.Node) *node.Position {
	return &node.Position{t.StartLine, getListPosEndLine(list)}
}

func NewNodeNodeListPosition(n node.Node, list []node.Node) *node.Position {
	return &node.Position{getNodePosStartLine(n), getListPosEndLine(list)}
}

func NewOptionalListTokensPosition(list []node.Node, t token.Token, endToken token.Token) *node.Position {
	if list == nil {
		return &node.Position{t.StartLine, endToken.EndLine}
	} else {
		return &node.Position{getListPosStartLine(list), endToken.EndLine}
	}
}

// AltIf Positions

func NewAltIfStartPosition(startToken token.Token) *node.Position {
	return &node.Position{startToken.StartLine, -1}
}
func NewAltIfPosition(startLine int, EndToken token.Token) *node.Position {
	return &node.Position{startLine, EndToken.EndLine}
}
