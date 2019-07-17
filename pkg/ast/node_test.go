package ast_test

import (
	"testing"

	"github.com/z7zmey/php-parser/pkg/ast"
	"gotest.tools/assert"
)

func TestNodeType(t *testing.T) {
	a := ast.NodeTypeAssignBitwiseOr

	assert.Assert(t, a.Is(ast.NodeClassTypeExpr))
	assert.Assert(t, a.Is(ast.NodeClassTypeAssign))
}

func TestNodeClassType(t *testing.T) {
	assert.Equal(t, ast.NodeClassTypeGeneral, ast.NodeClassType(1<<8))
	assert.Equal(t, ast.NodeClassTypeBinary, ast.NodeClassType(1<<14|1<<12))
}
