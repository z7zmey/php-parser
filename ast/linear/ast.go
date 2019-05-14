package linear

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/nested"
	"github.com/z7zmey/php-parser/scanner"
)

type lastTknCache struct {
	nodeID  NodeID
	tokenID TokenID
}

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	Tokens    *TokenStorage
	RootNode  NodeID

	lastTknCache lastTknCache
}

func (a *AST) Reset() {
	a.FileData = a.FileData[:0]
	a.Positions.Reset()
	a.Nodes.Reset()
	a.RootNode = 0
}

func (a *AST) Children(prevNodeID NodeID, parentNodeID NodeID, edgeType ast.EdgeType, children ...NodeID) NodeID {
	for _, childNodeID := range children {
		if childNodeID == 0 {
			continue
		}

		if prevNodeID == 0 {
			a.linkChild(parentNodeID, childNodeID)
		} else {
			a.linkNext(prevNodeID, childNodeID)
		}

		a.linkParent(childNodeID, parentNodeID, edgeType)

		prevNodeID = childNodeID
	}

	return prevNodeID
}

func (a *AST) linkParent(childNodeID, parentNodeID NodeID, key ast.EdgeType) {
	childNode := a.Nodes.Get(childNodeID)
	childNode.Parent = parentNodeID
	childNode.Key = key
	a.Nodes.Save(childNodeID, childNode)
}

func (a *AST) linkChild(parentNodeID, childNodeID NodeID) {
	parentNode := a.Nodes.Get(parentNodeID)
	parentNode.Child = childNodeID
	a.Nodes.Save(parentNodeID, parentNode)
}

func (a *AST) linkNext(prevNodeID, nextNodeID NodeID) {
	prevNode := a.Nodes.Get(prevNodeID)
	prevNode.Next = nextNodeID
	a.Nodes.Save(prevNodeID, prevNode)
}

func (a *AST) lastToken(nodeID NodeID) TokenID {
	if a.lastTknCache.nodeID == nodeID {
		return a.lastTknCache.tokenID
	}

	node := a.Nodes.Get(nodeID)
	tknID := node.Tkn

	if tknID == 0 {
		return tknID
	}

	for {
		token := a.Tokens.Get(tknID)

		if token.Next == 0 {
			break
		}

		tknID = token.Next
	}

	a.lastTknCache = lastTknCache{
		nodeID:  nodeID,
		tokenID: tknID,
	}

	return tknID
}

func (a *AST) AppendTokens(nodeID NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	lastTokenID := a.lastToken(nodeID)

	for _, str := range ffStrs {
		tkn := a.convertToken(str)
		tkn.Group = group
		tokenID := a.Tokens.Create(tkn)

		if lastTokenID == 0 {
			node := a.Nodes.Get(nodeID)
			node.Tkn = tokenID
			a.Nodes.Save(nodeID, node)
		} else {
			prevString := a.Tokens.Get(lastTokenID)
			prevString.Next = tokenID
			a.Tokens.Save(lastTokenID, prevString)
		}

		lastTokenID = tokenID
	}

	a.lastTknCache = lastTknCache{
		nodeID:  nodeID,
		tokenID: lastTokenID,
	}
}

func (a *AST) PrependTokens(nodeID NodeID, group ast.TokenGroup, ffStrs []scanner.Token) {
	node := a.Nodes.Get(nodeID)
	firstTokenID := node.Tkn

	var prevTokenID TokenID
	for _, str := range ffStrs {
		tkn := a.convertToken(str)
		tkn.Group = group
		tkn.Next = firstTokenID
		tokenID := a.Tokens.Create(tkn)

		if prevTokenID == 0 {
			node := a.Nodes.Get(nodeID)
			node.Tkn = tokenID
			a.Nodes.Save(nodeID, node)
		} else {
			prevToken := a.Tokens.Get(prevTokenID)
			prevToken.Next = tokenID
			a.Tokens.Save(prevTokenID, prevToken)
		}

		prevTokenID = tokenID
	}
}

func (a *AST) MoveStartTokens(src NodeID, dst NodeID) {
	srcNode := a.Nodes.Get(src)

	if srcNode.Tkn == 0 {
		return
	}

	srcStartFirstTkn := a.Tokens.Get(srcNode.Tkn)
	if srcStartFirstTkn.Group != ast.TokenGroupStart {
		return
	}

	srcStartLastTknID := srcNode.Tkn
	srcStartLastTkn := srcStartFirstTkn

	for {
		if srcStartLastTkn.Next == 0 {
			break
		}

		tkn := a.Tokens.Get(srcStartLastTkn.Next)
		if tkn.Group == ast.TokenGroupStart {
			srcStartLastTknID = srcStartLastTkn.Next
			srcStartLastTkn = tkn
		} else {
			break
		}
	}

	dstNode := a.Nodes.Get(dst)

	// move

	dstNode.Tkn, srcNode.Tkn, srcStartLastTkn.Next = srcNode.Tkn, srcStartLastTkn.Next, dstNode.Tkn

	// save

	a.Nodes.Save(src, srcNode)
	a.Nodes.Save(dst, dstNode)
	a.Tokens.Save(srcStartLastTknID, srcStartLastTkn)
}

func (a *AST) convertToken(token scanner.Token) Token {
	pos := ast.Position{
		PS: token.StartPos,
		PE: token.EndPos,
		LS: token.StartLine,
	}
	posID := a.Positions.Create(pos)

	return Token{
		Type: token.Type,
		Pos:  posID,
	}
}

func (a *AST) getListStartPosID(l []NodeID) PositionID {
	if len(l) > 0 {
		if l[0] == 0 {
			return 0
		}
		return a.Nodes.Get(l[0]).Pos
	}

	return 0
}

func (a *AST) getListEndPosID(l []NodeID) PositionID {
	if len(l) > 0 {
		if l[len(l)-1] == 0 {
			return 0
		}
		return a.Nodes.Get(l[len(l)-1]).Pos
	}

	return 0
}

func (a *AST) NewNodeListPosition(list []NodeID) PositionID {
	sPosID := a.getListStartPosID(list)
	ePosID := a.getListEndPosID(list)

	if sPosID == 0 || ePosID == 0 {
		return 0
	}

	s := a.Positions.Get(sPosID)
	e := a.Positions.Get(ePosID)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (a *AST) NewTokenPosition(t *scanner.Token) PositionID {
	if t == nil {
		return PositionID(0)
	}

	return a.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: t.EndPos,
		LS: t.StartLine,
		LE: t.EndLine,
	})
}

func (a *AST) NewTokensPosition(startToken *scanner.Token, endToken *scanner.Token) PositionID {
	if startToken == nil || endToken == nil {
		return PositionID(0)
	}

	return a.Positions.Create(ast.Position{
		PS: startToken.StartPos,
		PE: endToken.EndPos,
		LS: startToken.StartLine,
		LE: endToken.EndLine,
	})
}

func (a *AST) NewTokenNodePosition(t *scanner.Token, n NodeID) PositionID {
	if t == nil || n == 0 {
		return PositionID(0)
	}

	nPos := a.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	e := a.Positions.Get(nPos)

	return a.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (a *AST) NewNodeTokenPosition(n NodeID, t *scanner.Token) PositionID {
	if n == 0 || t == nil {
		return PositionID(0)
	}

	nPos := a.Nodes.Get(n).Pos
	if nPos == 0 {
		return 0
	}
	s := a.Positions.Get(nPos)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (a *AST) NewNodesPosition(startNodeID NodeID, endNodeID NodeID) PositionID {
	if startNodeID == 0 || endNodeID == 0 {
		return PositionID(0)
	}

	sPos := a.Nodes.Get(startNodeID).Pos
	ePos := a.Nodes.Get(endNodeID).Pos

	if sPos == 0 || ePos == 0 {
		return 0
	}
	s := a.Positions.Get(sPos)
	e := a.Positions.Get(ePos)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

// NewNodePosition returns new Position
func (a *AST) NewNodePosition(nodeID NodeID) PositionID {
	if nodeID == 0 {
		return PositionID(0)
	}

	posID := a.Nodes.Get(nodeID).Pos
	pos := a.Positions.Get(posID)

	return a.Positions.Create(ast.Position{
		PS: pos.PS,
		PE: pos.PE,
		LS: pos.LS,
		LE: pos.LE,
	})
}

func (a *AST) NewNodeListTokenPosition(list []NodeID, t *scanner.Token) PositionID {
	if list == nil || t == nil {
		return PositionID(0)
	}

	sPosID := a.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := a.Positions.Get(sPosID)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: t.EndPos,
		LS: s.LS,
		LE: t.EndLine,
	})
}

func (a *AST) NewTokenNodeListPosition(t *scanner.Token, list []NodeID) PositionID {
	if t == nil || list == nil {
		return PositionID(0)
	}

	ePosID := a.getListEndPosID(list)
	if ePosID == 0 {
		return 0
	}
	e := a.Positions.Get(ePosID)

	return a.Positions.Create(ast.Position{
		PS: t.StartPos,
		PE: e.PE,
		LS: t.StartLine,
		LE: e.LE,
	})
}

func (a *AST) NewNodeNodeListPosition(n NodeID, list []NodeID) PositionID {
	if n == 0 || list == nil {
		return PositionID(0)
	}

	nPos := a.Nodes.Get(n).Pos
	ePosID := a.getListEndPosID(list)
	if nPos == 0 || ePosID == 0 {
		return 0
	}
	s := a.Positions.Get(nPos)
	e := a.Positions.Get(ePosID)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (a *AST) NewNodeListNodePosition(list []NodeID, n NodeID) PositionID {
	if list == nil || n == 0 {
		return PositionID(0)
	}

	sPosID := a.getListStartPosID(list)
	nPos := a.Nodes.Get(n).Pos
	if sPosID == 0 || nPos == 0 {
		return 0
	}
	s := a.Positions.Get(sPosID)
	e := a.Positions.Get(nPos)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: e.PE,
		LS: s.LS,
		LE: e.LE,
	})
}

func (a *AST) NewOptionalListTokensPosition(list []NodeID, startToken *scanner.Token, endToken *scanner.Token) PositionID {
	if list == nil {
		if startToken == nil || endToken == nil {
			return PositionID(0)
		}

		return a.Positions.Create(ast.Position{
			PS: startToken.StartPos,
			PE: endToken.EndPos,
			LS: startToken.StartLine,
			LE: endToken.EndLine,
		})
	}

	if list == nil || endToken == nil {
		return PositionID(0)
	}

	sPosID := a.getListStartPosID(list)
	if sPosID == 0 {
		return 0
	}
	s := a.Positions.Get(sPosID)

	return a.Positions.Create(ast.Position{
		PS: s.PS,
		PE: endToken.EndPos,
		LS: s.LS,
		LE: endToken.EndLine,
	})
}

func (stxtree *AST) Traverse(v Visitor) {
	depth := 0
	curNodeID := stxtree.RootNode

	for {
		if curNodeID == 0 {
			break
		}

		curNode := stxtree.Nodes.Get(curNodeID)
		visitChild := v.VisitNode(stxtree, curNode, depth)

		if visitChild && curNode.Child != 0 {
			curNodeID = curNode.Child
			depth++
			continue
		}

		if curNode.Next != 0 {
			curNodeID = curNode.Next
			continue
		}

		for {
			if curNode.Parent == 0 {
				curNodeID = 0
				break
			}

			curNode = stxtree.Nodes.Get(curNode.Parent)
			depth--

			if curNode.Next != 0 {
				curNodeID = curNode.Next
				break
			}
		}
	}
}

func (stxtree *AST) Nested() nested.Node {
	depth := 0
	curNodeID := stxtree.RootNode

	stack := []nested.Node{}

	for {
		if curNodeID == 0 {
			break
		}

		curNode := stxtree.Nodes.Get(curNodeID)

		if len(stack) <= depth+1 {
			stack = append(stack, nested.Node{})
		}

		pos := stxtree.Positions.Get(curNode.Pos)

		stack[depth] = nested.Node{
			Type:     curNode.Type,
			Flags:    curNode.Flag,
			Position: pos,
			Tokens:   make(map[ast.TokenGroup][]nested.Token),
			Children: make(map[ast.EdgeType][]nested.Node),
		}

		tokenID := curNode.Tkn
		for {
			if tokenID == 0 {
				break
			}

			token := stxtree.Tokens.Get(tokenID)
			tokenPos := stxtree.Positions.Get(token.Pos)

			nestedToken := nested.Token{
				Type:  token.Type,
				Value: string(stxtree.FileData[tokenPos.PS:tokenPos.PE]),
			}

			stack[depth].Tokens[token.Group] = append(stack[depth].Tokens[token.Group], nestedToken)

			tokenID = token.Next
		}

		if curNode.Type.Is(ast.NodeClassTypeValue) {
			stack[depth].Value = string(stxtree.FileData[pos.PS:pos.PE])
		}

		if depth > 0 {
			stack[depth-1].Children[curNode.Key] = append(stack[depth-1].Children[curNode.Key], stack[depth])
		}

		if curNode.Child != 0 {
			curNodeID = curNode.Child
			depth++
			continue
		}

		if curNode.Next != 0 {
			curNodeID = curNode.Next
			continue
		}

		for {
			if curNode.Parent == 0 {
				curNodeID = 0
				break
			}

			curNode = stxtree.Nodes.Get(curNode.Parent)
			depth--

			if curNode.Next != 0 {
				curNodeID = curNode.Next
				break
			}
		}
	}

	return stack[0]
}
