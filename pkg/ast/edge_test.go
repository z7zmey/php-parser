package ast_test

import (
	"testing"

	"github.com/z7zmey/php-parser/pkg/ast"
	"gotest.tools/assert"
)

func TestEdgeClassType(t *testing.T) {
	e := ast.EdgeTypeConsts

	assert.Assert(t, e.Is(ast.EdgeClassTypeMultiple))
}
