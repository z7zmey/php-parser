package tree_test

import (
	"testing"

	"github.com/z7zmey/php-parser/internal/scanner"
	"github.com/z7zmey/php-parser/internal/tree"
	"github.com/z7zmey/php-parser/pkg/ast"
	"gotest.tools/assert"
)

type testVisitor struct {
	result []ast.NodeType
}

func (v *testVisitor) VisitNode(n ast.Node, depth int) bool {
	v.result = append(v.result, n.Type)
	return true
}

func TestAddTokens(t *testing.T) {
	stxtree := new(tree.Tree)

	tokens := stxtree.AddTokens([]ast.Token{
		{Type: scanner.T_OPEN_TAG},
		{Type: scanner.T_WHITESPACE},
	})

	expected := []ast.Token{
		{Type: scanner.T_OPEN_TAG},
		{Type: scanner.T_WHITESPACE},
	}

	assert.DeepEqual(t, expected, tokens)

	tokens = stxtree.AddTokens([]ast.Token{
		{Type: scanner.T_COMMENT},
		{Type: scanner.T_WHITESPACE},
	})

	expected = []ast.Token{
		{Type: scanner.T_COMMENT},
		{Type: scanner.T_WHITESPACE},
	}

	assert.DeepEqual(t, expected, tokens)

	tokens = stxtree.AppendTokens(len(tokens), []ast.Token{
		{Type: scanner.T_FUNCTION},
	})

	expected = []ast.Token{
		{Type: scanner.T_COMMENT},
		{Type: scanner.T_WHITESPACE},
		{Type: scanner.T_FUNCTION},
	}

	assert.DeepEqual(t, expected, tokens)
}

func TestGraphTraverseDFS(t *testing.T) {
	stxtree := new(tree.Tree)

	children := stxtree.AddNodes([]ast.Node{
		{Type: ast.NodeTypeNameNamePart, Value: []byte("foo"), Group: ast.NodeGroupParts},
		{Type: ast.NodeTypeNameNamePart, Value: []byte("bar"), Group: ast.NodeGroupParts},
	})

	children = stxtree.AddNodes([]ast.Node{
		{Type: ast.NodeTypeNameName, Group: ast.NodeGroupFunction, Children: children},
		{Type: ast.NodeTypeArgumentList, Group: ast.NodeGroupArgumentList},
	})

	children = stxtree.AddNodes([]ast.Node{
		{Type: ast.NodeTypeExprFunctionCall, Children: children},
	})

	expected := []ast.NodeType{
		ast.NodeTypeExprFunctionCall,
		ast.NodeTypeNameName,
		ast.NodeTypeNameNamePart,
		ast.NodeTypeNameNamePart,
		ast.NodeTypeArgumentList,
	}
	v := &testVisitor{}

	stxtree.Traverse(v)

	assert.DeepEqual(t, expected, v.result)
}
