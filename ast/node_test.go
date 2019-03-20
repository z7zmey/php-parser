package ast

import (
	"gotest.tools/assert"
	"testing"
)

func TestNodeType(t *testing.T) {
	a := NodeTypeAssignBitwiseOr

	assert.Assert(t, a.Is(NodeClassTypeExpr))
	assert.Assert(t, a.Is(NodeClassTypeAssign))
}

func TestNodeClassType(t *testing.T) {
	assert.Equal(t, NodeClassTypeGeneral, NodeClassType(1<<8))
	assert.Equal(t, NodeClassTypeBinary, NodeClassType(1<<14|1<<12))
}
