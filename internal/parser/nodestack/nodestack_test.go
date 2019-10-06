package nodestack

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/pkg/ast"
)

func TestNodeStack(t *testing.T) {
	s := NodeStack{}

	s.Push()

	s.Add(
		ast.Node{},
		ast.Node{},
		ast.Node{},
	)

	expected := []ast.Node{
		{Group: ast.NodeGroupArguments},
		{Group: ast.NodeGroupArguments},
		{Group: ast.NodeGroupArguments},
	}
	actual := s.Pop(ast.NodeGroupArguments)
	assert.DeepEqual(t, expected, actual)

	expected = []ast.Node{}
	actual = s.Pop(ast.NodeGroupArguments)
	assert.DeepEqual(t, expected, actual)
}

func TestNodeStackNested(t *testing.T) {
	s := NodeStack{}

	s.Push()

	s.Add(ast.Node{})
	s.Add(ast.Node{})
	s.Add(ast.Node{})

	s.Push(
		ast.Node{},
		ast.Node{},
	)

	expected := []ast.Node{
		{Group: ast.NodeGroupArguments},
		{Group: ast.NodeGroupArguments},
	}
	actual := s.Pop(ast.NodeGroupArguments)
	assert.DeepEqual(t, expected, actual)

	expected = []ast.Node{
		{Group: ast.NodeGroupStmts},
		{Group: ast.NodeGroupStmts},
		{Group: ast.NodeGroupStmts},
	}
	actual = s.Pop(ast.NodeGroupStmts)
	assert.DeepEqual(t, expected, actual)

	expected = []ast.Node{}
	actual = s.Pop()
	assert.DeepEqual(t, expected, actual)
}

func TestNodeStackMulti(t *testing.T) {
	s := NodeStack{}

	s.Push()

	s.Add(ast.Node{})
	s.Add(ast.Node{})
	s.Add(ast.Node{})

	s.Push()
	s.Add(ast.Node{})
	s.Add(ast.Node{})

	expected := []ast.Node{
		{Group: ast.NodeGroupStmts},
		{Group: ast.NodeGroupStmts},
		{Group: ast.NodeGroupStmts},
		{Group: ast.NodeGroupArguments},
		{Group: ast.NodeGroupArguments},
	}
	actual := s.Pop(ast.NodeGroupArguments, ast.NodeGroupStmts)
	assert.DeepEqual(t, expected, actual)
}
