package linear

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/scanner"
)

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	RootNode  NodeID
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

func (ast *AST) Traverse(v Visitor) {
	depth := 0
	curNodeID := ast.RootNode

	for {
		if curNodeID == 0 {
			break
		}

		curNode := ast.Nodes.Get(curNodeID)
		visitChild := v.VisitNode(curNode, depth)

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

			curNode = ast.Nodes.Get(curNode.Parent)
			depth--

			if curNode.Next != 0 {
				curNodeID = curNode.Next
				break
			}
		}
	}
}
