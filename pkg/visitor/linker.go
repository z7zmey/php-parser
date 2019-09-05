package visitor

import "github.com/z7zmey/php-parser/pkg/ast"

type Linker struct {
	stack []*ast.Node
}

func (v *Linker) VisitNode(n ast.SimpleNode, depth int) bool {
	if len(v.stack) == depth {
		v.stack = append(v.stack, nil)
	}

	v.stack[depth] = &ast.Node{
		SimpleNode: n,
		Tokens:     make(map[ast.TokenGroup][]ast.Token),
	}

	if depth == 0 {
		return true
	}

	parent := v.stack[depth-1]

	if parent.Children == nil {
		parent.Children = make(map[ast.NodeGroup][]*ast.Node)
	}

	parent.Children[n.Group] = append(parent.Children[n.Group], v.stack[depth])

	return true
}

func (v *Linker) VisitPosition(p ast.Position, depth int) bool {
	v.stack[depth-1].Position = p

	return true
}

func (v *Linker) VisitToken(t ast.Token, depth int) bool {
	v.stack[depth-1].Tokens[t.Group] = append(v.stack[depth].Tokens[t.Group], t)

	return true
}

func (v *Linker) GetRoot() ast.Node {
	return *v.stack[0]
}
