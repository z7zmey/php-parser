package linear

import (
	"github.com/z7zmey/php-parser/ast"
	"github.com/z7zmey/php-parser/ast/nested"
)

type AST struct {
	FileData  []byte
	Positions *PositionStorage
	Nodes     *NodeStorage
	Tokens    *TokenStorage
	RootNode  NodeID
}

func (a *AST) Reset() {
	a.FileData = a.FileData[:0]
	a.Positions.Reset()
	a.Nodes.Reset()
	a.RootNode = 0
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
