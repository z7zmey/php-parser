package php7

import (
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/ast"
)

func TestNodeType(t *testing.T) {
	s := stackedNodeList{}

	s.push()

	s.add(ast.NodeID(1))
	s.add(ast.NodeID(2))
	s.add(ast.NodeID(3))

	s.push()
	s.add(ast.NodeID(4))
	s.add(ast.NodeID(5))

	expected := []ast.NodeID{ast.NodeID(4), ast.NodeID(5)}
	actual := s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []ast.NodeID{ast.NodeID(1), ast.NodeID(2), ast.NodeID(3)}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)

	expected = []ast.NodeID{}
	actual = s.pop()
	assert.DeepEqual(t, expected, actual)
}
