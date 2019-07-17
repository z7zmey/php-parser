package visitor

import "github.com/z7zmey/php-parser/pkg/ast"

type Linker struct {
	stack []*ast.Node
}

func (v *Linker) VisitNode(n ast.Node, group ast.NodeGroup, depth int) bool {
	if len(v.stack) == depth {
		v.stack = append(v.stack, &n)
	}

	v.stack[depth] = &n

	if depth == 0 {
		return true
	}

	parent := v.stack[depth-1]

	if parent.Children == nil {
		parent.Children = make(map[ast.NodeGroup][]*ast.Node)
	}

	parent.Children[group] = append(parent.Children[group], &n)

	return true
}

func (v *Linker) GetRoot() ast.Node {
	return *v.stack[0]
}
