package parser

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

type startPos struct {
	startLine int
	startPos  int
}

type endPos struct {
	endLine int
	endPos  int
}

func getListStartPos(l []node.Node) startPos {
	if l == nil {
		return startPos{-1, -1}
	}

	if len(l) == 0 {
		return startPos{-1, -1}
	}

	return getNodeStartPos(l[0])
}

func getNodeStartPos(n node.Node) startPos {
	sl := -1
	sp := -1

	if n == nil {
		return startPos{-1, -1}
	}

	p := n.Position()
	if p != nil {
		sl = p.StartLine
		sp = p.StartPos
	}

	return startPos{sl, sp}
}

func getListEndPos(l []node.Node) endPos {
	if l == nil {
		return endPos{-1, -1}
	}

	if len(l) == 0 {
		return endPos{-1, -1}
	}

	return getNodeEndPos(l[len(l)-1])
}

func getNodeEndPos(n node.Node) endPos {
	el := -1
	ep := -1

	if n == nil {
		return endPos{-1, -1}
	}

	p := n.Position()
	if p != nil {
		el = p.EndLine
		ep = p.EndPos
	}

	return endPos{el, ep}
}

func NewNodeListPosition(list []node.Node) *node.Position {
	return &node.Position{
		getListStartPos(list).startLine,
		getListEndPos(list).endLine,
		getListStartPos(list).startPos,
		getListEndPos(list).endPos,
	}
}

func NewNodePosition(n node.Node) *node.Position {
	return &node.Position{
		getNodeStartPos(n).startLine,
		getNodeEndPos(n).endLine,
		getNodeStartPos(n).startPos,
		getNodeEndPos(n).endPos,
	}
}

func NewTokenPosition(t token.Token) *node.Position {
	return &node.Position{
		t.StartLine,
		t.EndLine,
		t.StartPos,
		t.EndPos,
	}
}

func NewTokensPosition(startToken token.Token, endToken token.Token) *node.Position {
	return &node.Position{
		startToken.StartLine,
		endToken.EndLine,
		startToken.StartPos,
		endToken.EndPos,
	}
}

func NewTokenNodePosition(t token.Token, n node.Node) *node.Position {
	return &node.Position{
		t.StartLine,
		getNodeEndPos(n).endLine,
		t.StartPos,
		getNodeEndPos(n).endPos,
	}
}

func NewNodeTokenPosition(n node.Node, t token.Token) *node.Position {
	return &node.Position{
		getNodeStartPos(n).startLine,
		t.EndLine,
		getNodeStartPos(n).startPos,
		t.EndPos,
	}
}

func NewNodesPosition(startNode node.Node, endNode node.Node) *node.Position {
	return &node.Position{
		getNodeStartPos(startNode).startLine,
		getNodeEndPos(endNode).endLine,
		getNodeStartPos(startNode).startPos,
		getNodeEndPos(endNode).endPos,
	}
}

func NewNodeListTokenPosition(list []node.Node, t token.Token) *node.Position {
	return &node.Position{
		getListStartPos(list).startLine,
		t.EndLine,
		getListStartPos(list).startPos,
		t.EndPos,
	}
}

func NewTokenNodeListPosition(t token.Token, list []node.Node) *node.Position {
	return &node.Position{
		t.StartLine,
		getListEndPos(list).endLine,
		t.StartPos,
		getListEndPos(list).endPos,
	}
}

func NewNodeNodeListPosition(n node.Node, list []node.Node) *node.Position {
	return &node.Position{
		getNodeStartPos(n).startLine,
		getListEndPos(list).endLine,
		getNodeStartPos(n).startPos,
		getListEndPos(list).endPos,
	}
}

func NewOptionalListTokensPosition(list []node.Node, t token.Token, endToken token.Token) *node.Position {
	if list == nil {
		return &node.Position{
			t.StartLine,
			endToken.EndLine,
			t.StartPos,
			endToken.EndPos,
		}
	} else {
		return &node.Position{
			getListStartPos(list).startLine,
			endToken.EndLine,
			getListStartPos(list).startPos,
			endToken.EndPos,
		}
	}
}

func ListGetFirstNodeComments(list []node.Node) []comment.Comment {
	if len(list) == 0 {
		return nil
	}

	node := list[0]

	return *node.Comments()
}
